package database

import (
	"context"

	"github.com/barealek/chatapp/types"
)

type Database interface {
	Disconnect(ctx context.Context) error

	SaveItem(ctx context.Context, item any) error
	DeleteItem(ctx context.Context, item any) error

	// sessions
	GetSessionFromID(ctx context.Context, sessid string) (types.Session, error)

	GetUserFromID(ctx context.Context, id string) (types.User, error)
	GetUserFromName(ctx context.Context, name string) (types.User, error)
}
