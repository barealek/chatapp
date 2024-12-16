package api

import (
	"encoding/json"
	"net/http"

	"github.com/barealek/chatapp/middleware"
	"github.com/barealek/chatapp/types"
	"github.com/charmbracelet/log"
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

func (a *Api) handlerGetMessages(w http.ResponseWriter, r *http.Request) {
	msgs, err := a.db.GetLatestMessages(r.Context(), 50)
	if err != nil {

		http.Error(w, "error fetching messages", http.StatusInternalServerError)
		log.Error("error fetching messages", "error", err)
		return
	}

	// Sort messages after the msg.SentAt int
	// to make sure the messages are in the right order
	sortMessages(msgs)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(msgs)
}

func sortMessages(msgs []types.Message) {
	// Bubble sort
	for i := 0; i < len(msgs); i++ {
		for j := 0; j < len(msgs)-i-1; j++ {
			if msgs[j].SentAt > msgs[j+1].SentAt {
				msgs[j], msgs[j+1] = msgs[j+1], msgs[j]
			}
		}
	}
}
