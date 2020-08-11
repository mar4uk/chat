package http

import (
	"encoding/json"
	"net/http"

	"github.com/mar4uk/chat/internal/ctxutils"

	"github.com/mar4uk/chat/internal/app"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-chi/render"
)

// Message is
type Message struct {
	ID        primitive.ObjectID `json:"id"`
	UserID    int64              `json:"userId"`
	Text      string             `json:"text"`
	CreatedAt string             `json:"createdAt"`
}

type getMessagesHandler struct {
	app app.App
}

type createMessageHandler struct {
	app app.App
}

func (h *getMessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chat := ctxutils.GetChat(ctx)

	appMessages, err := h.app.GetMessages(ctx, chat.ID)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	messages := make([]*Message, len(appMessages))
	for i, message := range appMessages {
		messages[i] = &Message{
			ID:        message.ID,
			UserID:    message.UserID,
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		}
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, MessagesListResponse(messages))
}

func (h *createMessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var m *Message

	ctx := r.Context()
	chat := ctxutils.GetChat(ctx)

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	err := h.app.CreateMessage(ctx, chat.ID, app.Message{
		ID:        m.ID,
		UserID:    m.UserID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
	})
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, m)
}
