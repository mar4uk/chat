package app

import (
	"context"
	"time"

	"github.com/mar4uk/chat/internal/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorName string

const (
	RecordNotFound ErrorName = "Record not found"
	InvalidArgs    ErrorName = "Invalid params"
)

// App is interface of chat application
type App interface {
	GetMessageByID(ctx context.Context, messageID primitive.ObjectID) (Message, *Error)
	GetMessages(ctx context.Context, chatID int64) ([]*Message, *Error)
	GetChat(ctx context.Context, chatID int64) (*Chat, *Error)
	CreateMessage(ctx context.Context, chatID int64, message Message) (primitive.ObjectID, *Error)
}

type app struct {
	db store.Database
}

// Error is struc for app error
type Error struct {
	Error error
	Name  ErrorName
}

// User is struct in Message struct
type User struct {
	ID    primitive.ObjectID
	Name  string
	Email string
}

// Message is
type Message struct {
	ID        primitive.ObjectID
	User      User
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

func (a *app) GetMessages(ctx context.Context, chatID int64) ([]*Message, *Error) {
	dbMessages, err := a.db.GetMessages(ctx, chatID)

	if err != nil {
		return nil, &Error{err, RecordNotFound}
	}

	var messages []*Message
	users := make(map[primitive.ObjectID]User)
	usersIDs := make(map[primitive.ObjectID]bool)

	for _, message := range dbMessages {
		usersIDs[message.UserID] = true
	}

	for userID := range usersIDs {
		dbUser, err := a.db.GetUserByID(ctx, userID)

		if err != nil {
			return nil, &Error{err, RecordNotFound}
		}

		users[userID] = User{
			ID:    userID,
			Name:  dbUser.Name,
			Email: dbUser.Email,
		}
	}

	for _, message := range dbMessages {
		messages = append(messages, &Message{
			ID:        message.ID,
			User:      users[message.UserID],
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		})
	}

	return messages, nil
}

func (a *app) GetMessageByID(ctx context.Context, messageID primitive.ObjectID) (Message, *Error) {
	dbMessage, err := a.db.GetMessageByID(ctx, messageID)

	if err != nil {
		return Message{}, &Error{err, RecordNotFound}
	}

	dbUser, err := a.db.GetUserByID(ctx, dbMessage.UserID)

	if err != nil {
		return Message{}, &Error{err, RecordNotFound}
	}

	return Message{
		ID: dbMessage.ID,
		User: User{
			ID:    dbUser.ID,
			Name:  dbUser.Name,
			Email: dbUser.Email,
		},
		Text:      dbMessage.Text,
		CreatedAt: dbMessage.CreatedAt,
	}, nil
}

func (a *app) GetChat(ctx context.Context, chatID int64) (*Chat, *Error) {
	dbChat, err := a.db.GetChat(ctx, chatID)

	if err != nil {
		return nil, &Error{err, RecordNotFound}
	}

	return &Chat{
		ID:    dbChat.ID,
		Title: dbChat.Title,
	}, nil
}

func (a *app) CreateMessage(ctx context.Context, chatID int64, message Message) (primitive.ObjectID, *Error) {
	messageID, err := a.db.CreateMessage(ctx, store.Message{
		ChatID:    chatID,
		UserID:    message.User.ID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	})

	if err != nil {
		return messageID, &Error{err, InvalidArgs}
	}

	return messageID, nil
}
