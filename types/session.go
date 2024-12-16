package types

import (
	"time"
)

type Session struct {
	ID          string        `bson:"id"`
	UserID      string        `bson:"userid"`
	CreatedAt   time.Time     `bson:"created_at"`
	CreatedFrom string        `bson:"created_from"`
	TTL         time.Duration `bson:"ttl"`
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.CreatedAt.Add(s.TTL))
}
