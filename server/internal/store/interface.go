package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message is
type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ChatID    int64              `bson:"chatId"`
	UserID    primitive.ObjectID `bson:"userId"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"createdAt"`
}

// Chat is
type Chat struct {
	ID    int64
	Title string
}

// User is
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

// Database is main interface to any DB
type Database interface {
	GetMessages(ctx context.Context, chatID int64) ([]Message, error)
	GetMessageByID(ctx context.Context, id primitive.ObjectID) (Message, error)
	GetChat(ctx context.Context, chatID int64) (*Chat, error)
	CreateMessage(ctx context.Context, message Message) (primitive.ObjectID, error)
	CreateUser(ctx context.Context, user User) (primitive.ObjectID, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	Close(ctx context.Context)
}
