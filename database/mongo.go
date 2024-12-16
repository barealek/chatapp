package database

import (
	"context"
	"fmt"

	"github.com/barealek/chatapp/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(ctx context.Context, conn string) (Database, error) {
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}
	err = cl.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	inst := &mongodb{cl: cl, dbname: "dev"}
	return inst, err
}

type mongodb struct {
	cl     *mongo.Client
	dbname string
}

func (m *mongodb) Disconnect(ctx context.Context) error {
	return m.cl.Disconnect(ctx)
}

func (m *mongodb) GetSessionFromID(ctx context.Context, sessid string) (types.Session, error) {

	db := m.getDatabase()
	coll := db.Collection("sessions")

	query := bson.M{"id": sessid}
	var s types.Session
	err := coll.FindOne(ctx, query).Decode(&s)
	if err != nil {
		return types.Session{}, err
	}
	return s, nil
}

func (m *mongodb) SaveItem(ctx context.Context, item any) error {
	db := m.getDatabase()
	var collName string

	switch item.(type) {
	case types.User, *types.User:
		collName = "users"
	case types.Session:
		collName = "sessions"
	case types.Message:
		collName = "messages"
	default:
		panic(fmt.Sprintf("type %T is not key registered in the database", item))
	}

	coll := db.Collection(collName)
	_, err := coll.InsertOne(ctx, item)
	return err
}

func (m *mongodb) DeleteItem(ctx context.Context, item any) error {
	db := m.getDatabase()
	var collName string

	switch item.(type) {
	case types.User, *types.User:
		collName = "users"
	case types.Session:
		collName = "sessions"
	}

	if collName == "" {
		panic(fmt.Sprintf("type %T is not key registered in the database", item))
	}

	coll := db.Collection(collName)
	_, err := coll.DeleteOne(ctx, item)
	return err
}

func (m *mongodb) GetUserFromID(ctx context.Context, id string) (types.User, error) {
	db := m.getDatabase()
	coll := db.Collection("users")

	query := bson.M{"id": id}
	var u types.User
	err := coll.FindOne(ctx, query).Decode(&u)
	if err != nil {
		return types.User{}, err
	}
	return u, nil
}

func (m *mongodb) GetUserFromName(ctx context.Context, name string) (types.User, error) {

	db := m.getDatabase()
	coll := db.Collection("users")

	query := bson.M{"name": name}
	var u types.User
	err := coll.FindOne(ctx, query).Decode(&u)
	if err != nil {
		return types.User{}, err
	}
	return u, nil
}

func (m *mongodb) getDatabase() *mongo.Database {
	return m.cl.Database(m.dbname)
}

func (m *mongodb) GetLatestMessages(ctx context.Context, n int) ([]types.Message, error) {
	db := m.getDatabase()
	coll := db.Collection("messages")

	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}).SetLimit(int64(n))
	cursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var messages []types.Message
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
