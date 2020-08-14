package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mar4uk/chat/internal/ctxutils"

	"github.com/mar4uk/chat/internal/app"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-chi/render"
)

// Message is
type Message struct {
	ID        primitive.ObjectID `json:"id"`
	ChatID    int64              `json:"chatId"`
	UserID    int64              `json:"userId"`
	Text      string             `json:"text"`
	CreatedAt time.Time          `json:"createdAt"`
}

type getMessagesHandler struct {
	app app.App
}

type createMessageHandler struct {
	app app.App
}

type websocketHandler struct {
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

	_, err := h.app.CreateMessage(ctx, chat.ID, app.Message{
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}

		var m *Message

		if err := json.Unmarshal(message, &m); err != nil {
			fmt.Println("unmarshal:", err)
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		messageID, err := h.app.CreateMessage(r.Context(), m.ChatID, app.Message{
			UserID:    m.UserID,
			Text:      m.Text,
			CreatedAt: m.CreatedAt,
		})

		m.ID = messageID
		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		msg, err := json.Marshal(m)

		if err != nil {
			fmt.Println("marshal:", err)
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		fmt.Printf("recv: %s", msg)

		err = c.WriteMessage(mt, msg)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}
