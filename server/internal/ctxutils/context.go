package ctxutils

import (
	"context"

	"github.com/mar4uk/chat/internal/app"
)

type contextID int

const (
	chatID contextID = iota
	userID contextID = iota
)

// GetChat gets chat
func GetChat(ctx context.Context) *app.Chat {
	if ret, ok := ctx.Value(chatID).(*app.Chat); !ok {
		panic("context without chat")
	} else {
		return ret
	}
}

// SetChat sets chat
func SetChat(ctx context.Context, chat *app.Chat) context.Context {
	return context.WithValue(ctx, chatID, chat)
}

// GetUser gets user
func GetUser(ctx context.Context) *app.User {
	if ret, ok := ctx.Value(userID).(*app.User); !ok {
		panic("context without user")
	} else {
		return ret
	}
}

// SetUser sets user
func SetUser(ctx context.Context, user *app.User) context.Context {
	return context.WithValue(ctx, userID, user)
}
