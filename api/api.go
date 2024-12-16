package api

import (
	"net/http"

	"github.com/barealek/chatapp/database"
	"github.com/barealek/chatapp/middleware"
	"github.com/gorilla/mux"
)

type Api struct {
	port int
	db   database.Database
}

func NewAPI(db database.Database, port int) http.Handler {
	a := &Api{
		port: port,
		db:   db,
	}
	return a.RegisterRoutes()
}

func (a *Api) RegisterRoutes() http.Handler {
	m := mux.NewRouter().PathPrefix("/api").Subrouter()

	auth := m.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", a.handlerLogin).Methods("POST")
	auth.HandleFunc("/signup", a.handlerRegisterUser).Methods("POST")
	auth.HandleFunc("/logout", a.handlerLogout).Methods("GET")

	sess := m.PathPrefix("/session").Subrouter()
	sess.Use(middleware.AuthMiddleware(a.db))
	sess.HandleFunc("", a.handlerSession).Methods("GET")

	messages := m.PathPrefix("/messages").Subrouter()
	messages.Use(middleware.AuthMiddleware(a.db))
	messages.HandleFunc("", a.handlerGetMessages).Methods("GET")
	messages.HandleFunc("", a.handlerSendMessage).Methods("POST")

	return m
}
