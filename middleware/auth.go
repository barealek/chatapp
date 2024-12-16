package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/barealek/chatapp/database"
	"github.com/barealek/chatapp/pkg/errorid"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthMiddleware(db database.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session")
			fmt.Printf("r.Header: %v\n", r.Header)
			if err != nil {
				if err == http.ErrNoCookie {
					http.Error(w, "not logged in", http.StatusUnauthorized)
					return
				}
				erid := errorid.GenErrorID()
				log.Error("error while retrieving session cookie", "error", err, "code", erid)
				internalservererror(w, "error decoding cookie", erid)
				return
			}

			session, err := db.GetSessionFromID(r.Context(), c.Value)
			if err == mongo.ErrNoDocuments {
				unauthorized(w)
				return
			}
			if err != nil {
				errid := errorid.GenErrorID()
				log.Error("error retrieving session", "sessionid", c.Value, "error", err, "code", errid)
				internalservererror(w, "error retrieving session", errid)
				return
			}

			if session.IsExpired() {
				// Expire session cookie
				http.SetCookie(w, &http.Cookie{
					Name:     "session",
					Domain:   "localhost",
					Value:    "/",
					Expires:  time.Unix(0, 0),
					HttpOnly: true,
				})
				http.Error(w, "your session has expired", http.StatusForbidden)
				return
			}

			user, err := db.GetUserFromID(r.Context(), session.UserID)
			if err != nil {
				errid := errorid.GenErrorID()
				log.Error("error retrieving session user", "sessionid", c.Value, "error", err, "code", errid)
				internalservererror(w, "error retrieving session user", errid)
				return
			}

			ctx := context.WithValue(r.Context(), ContextKeyUser, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
