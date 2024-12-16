package api

import (
	"encoding/json"
	"net/http"

	"github.com/barealek/chatapp/middleware"
	"github.com/barealek/chatapp/types"
)

func (a *Api) handlerSendMessage(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(middleware.ContextKeyUser).(types.User)

	type bodyData struct {
		Message string `json:"msg"`
	}

	var body bodyData
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if body.Message == "" {
		http.Error(w, "`msg` must not be empty", http.StatusBadRequest)
		return
	}

	msg := types.NewMessage(u, body.Message)

	a.db.SaveItem(r.Context(), msg)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "successful",
		"message": msg,
	})
}
