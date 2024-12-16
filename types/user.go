package types

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Password []byte `bson:"password" json:"-"`
	IsAdmin  bool   `bson:"is_admin" json:"is_admin"`
}

func (u *User) GenerateSession(ip string, d time.Duration) Session {
	var idToken string
	idToken += base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))
	idToken += base64.StdEncoding.EncodeToString([]byte(time.Now().String()))

	c := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var res string
	for range 8 {
		res += string(c[rand.Intn(len(c))])
	}
	idToken += base64.StdEncoding.EncodeToString([]byte(res))

	return Session{
		ID:          idToken,
		UserID:      u.ID,
		CreatedAt:   time.Now(),
		CreatedFrom: ip,
		TTL:         d,
	}
}

func NewUser(name string, passwordhash []byte) User {
	return User{
		Name:     name,
		ID:       uuid.NewString(),
		Password: passwordhash,
		IsAdmin:  false,
	}
}
