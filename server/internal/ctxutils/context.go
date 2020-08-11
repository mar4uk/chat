package ctxutils

import (
	"context"

	"github.com/mar4uk/chat/internal/app"
)

type contextID int

const (
	chatID contextID = iota
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
