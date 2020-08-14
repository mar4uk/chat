package app

import (
	"context"
	"time"

	"github.com/mar4uk/chat/internal/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// App is interface of chat application
type App interface {
	GetMessages(ctx context.Context, chatID int64) ([]*Message, error)
	GetChat(ctx context.Context, chatID int64) (*Chat, error)
	CreateMessage(ctx context.Context, chatID int64, message Message) (primitive.ObjectID, error)
}

type app struct {
	db store.Database
}

// Message is
type Message struct {
	ID        primitive.ObjectID
	UserID    int64
	Text      string
	CreatedAt time.Time
}

// Chat is
type Chat struct {
	ID    int64
	Title string
}

// NewApp is chat initialization function
func NewApp(db store.Database) App {
	return &app{
		db: db,
	}
}

func (a *app) GetMessages(ctx context.Context, chatID int64) ([]*Message, error) {
	dbMessages, err := a.db.GetMessages(ctx, chatID)

	if err != nil {
		return nil, err
	}

	var messages []*Message

	for _, message := range dbMessages {
		messages = append(messages, &Message{
			ID:        message.ID,
			UserID:    message.UserID,
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		})
	}

	return messages, nil
}

func (a *app) GetChat(ctx context.Context, chatID int64) (*Chat, error) {
	dbChat, err := a.db.GetChat(ctx, chatID)

	if err != nil {
		return nil, err
	}

	return &Chat{
		ID:    dbChat.ID,
		Title: dbChat.Title,
	}, nil
}

func (a *app) CreateMessage(ctx context.Context, chatID int64, message Message) (primitive.ObjectID, error) {
	messageID, err := a.db.CreateMessage(ctx, store.Message{
		ChatID:    chatID,
		UserID:    message.UserID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	})

	if err != nil {
		return messageID, err
	}

	return messageID, nil
}
