package server

import (
	"fmt"
	"net/http"

	"github.com/barealek/chatapp/api"
	"github.com/barealek/chatapp/database"
)

func NewServer(db database.Database, port int) *http.Server {
	bind := fmt.Sprintf(":%d", port)
	a := api.NewAPI(db, port)

	s := &http.Server{
		Handler: a,
		Addr:    bind,
	}

	return s
}
