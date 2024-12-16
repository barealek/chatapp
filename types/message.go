package types

import "time"

type Message struct {
	AuthorID   string `json:"author_id" bson:"author_id"`
	AuthorName string `json:"author_name" bson:"author_name"`
	Content    string `json:"content" bson:"content"`
	Likes      int    `json:"likes" bson:"likes"`
	SentAt     int64  `json:"sent_at" bson:"sent_at"`
}

func NewMessage(author User, content string) Message {
	return Message{
		AuthorID:   author.ID,
		AuthorName: author.Name,
		Content:    content,
		Likes:      0,
		SentAt:     time.Now().UnixMilli(),
	}
}
