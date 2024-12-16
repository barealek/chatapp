package api

import (
	"encoding/json"
	"net/http"

	"github.com/barealek/chatapp/middleware"
	"github.com/barealek/chatapp/types"
)

func (a *Api) handlerSession(w http.ResponseWriter, r *http.Request) {
	_u := r.Context().Value(middleware.ContextKeyUser)
	u, ok := _u.(types.User)
	if !ok {
		http.Error(w, "internal error casting user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}
