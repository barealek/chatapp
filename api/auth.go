package api

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/barealek/chatapp/pkg/hashing"
	"github.com/barealek/chatapp/types"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *Api) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type bodyData struct {
		Username string `json:"user"`
		Password string `json:"password"`
	}

	var body bodyData

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if body.Password == "" || body.Username == "" {
		http.Error(w, "missing required params", http.StatusBadRequest)
		return
	}

	hashedPwd := hashing.HashString(body.Password)
	u, err := a.db.GetUserFromName(r.Context(), body.Username)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			goto passwordhandle
		}
		http.Error(w, "error trying to log you in", http.StatusUnauthorized)
		log.Info("error trying to fetch user", "username", u.Name, "error", err)
		return
	}

passwordhandle:
	if !slices.Equal(u.Password, hashedPwd) || err == mongo.ErrNoDocuments {
		http.Error(w, "wrong username or password", http.StatusUnauthorized)
		log.Info("", "error", err)
		return
	}

	sess := u.GenerateSession(strings.Split(r.RemoteAddr, ":")[0], 1*time.Hour)

	err = a.db.SaveItem(r.Context(), sess)
	if err != nil {
		http.Error(w, "error trying to log you in", http.StatusInternalServerError)
		log.Info("error trying to create session", "user", u.Name, "userid", u.ID, "error", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sess.ID,
		Path:     "/api",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}

// handlerRegisterUser handles the user registration logic.
//
// Excepted body:
//
//	{
//	   user: string
//		password: string
//	}
func (a *Api) handlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	type bodydata struct {
		Username string `json:"user"`
		Password string `json:"password"`
	}

	var body bodydata

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "error decoding", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Check a user with the username doesn't exist
	_, err = a.db.GetUserFromName(r.Context(), body.Username)
	switch err {
	case mongo.ErrNoDocuments:
		u := types.NewUser(body.Username, hashing.HashString(body.Password))

		err = a.db.SaveItem(r.Context(), u)
		if err != nil {
			http.Error(w, "internal server error trying to create your account", http.StatusInternalServerError)
			log.Error("error creating user", "username", u.Name, "error", err)
			return
		}

		sess := u.GenerateSession(strings.Split(r.RemoteAddr, ":")[0], 1*time.Hour)
		a.db.SaveItem(r.Context(), sess)

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			SameSite: http.SameSiteStrictMode,
			Path:     "/api",
			Value:    sess.ID,
			HttpOnly: true,
		})

		json.NewEncoder(w).Encode(map[string]string{
			"status": "success",
		})

	case nil:
		http.Error(w, "user with that name already exists", http.StatusForbidden)
		return

	default:
		http.Error(w, "server error trying to create your account", http.StatusInternalServerError)
		log.Error("error trying to fetch preexisting user")
		return
	}
}

func (a *Api) handlerLogout(w http.ResponseWriter, r *http.Request) {

	sessCookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "error fetching cookie", http.StatusInternalServerError)
		return
	}

	sessId := sessCookie.Value

	sess, err := a.db.GetSessionFromID(r.Context(), sessId)
	if err != nil {
		http.Error(w, "not logged in", http.StatusUnauthorized)
	}

	err = a.db.DeleteItem(r.Context(), sess)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			panic("ingen dokumenter")
		}
		panic("fejl" + err.Error())
	}
	log.Info("user signed out", "session", sess.ID)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "/",
		Path:     "/api",
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"status": "successfully logged out",
	})
}
