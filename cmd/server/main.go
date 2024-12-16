package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/barealek/chatapp/database"
	"github.com/barealek/chatapp/pkg/must"
	"github.com/barealek/chatapp/server"
	"github.com/charmbracelet/log"
)

func main() {
	log.Info("initializing base context")
	ctx, finish := context.WithCancel(context.Background())

	port := flag.Int("p", 3000, "port to use when binding")
	flag.Parse()

	log.Info("creating server")

	db := must.Must(database.NewMongo(ctx, os.Getenv("DB_CONN")))

	s := server.NewServer(db, *port)

	go grace(ctx, s, finish)

	log.Info("success. listening...", "port", *port)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("error starting server", "error", err)
	}

	<-ctx.Done()

	log.Info("running post-shutdown jobs")
}

func grace(ctx context.Context, server *http.Server, finish context.CancelFunc) {
	signalctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	<-signalctx.Done()

	log.Info("shutting down")

	timeoutctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutctx); err != nil {
		log.Warn("server forced to shutdown", "error", err)
	}

	finish()
}
