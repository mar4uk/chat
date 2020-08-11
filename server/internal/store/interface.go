package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message is
type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ChatID    int64              `bson:"chatId"`
	UserID    int64              `bson:"userId"`
	Text      string             `bson:"text"`
	CreatedAt string             `bson:"createdAt"`
}

// Chat is
type Chat struct {
	ID    int64
	Title string
}

// Database is main interface to any DB
type Database interface {
	GetMessages(ctx context.Context, chatID int64) ([]Message, error)
	GetChat(ctx context.Context, chatID int64) (*Chat, error)
	CreateMessage(ctx context.Context, message Message) error
	Close(ctx context.Context)
}
