package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"github.com/mar4uk/chat/internal/app"
)

type websocketHandler struct {
	app app.App
}

func (h *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	channel, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer channel.Close()
	for {
		messageType, message, err := channel.ReadMessage()
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

		messageID, err := h.app.CreateMessage(ctx, m.ChatID, app.Message{
			User: app.User{
				ID: m.User.ID,
			},
			Text:      m.Text,
			CreatedAt: m.CreatedAt,
		})

		m.ID = messageID
		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		msg, err := h.app.GetMessageByID(ctx, messageID)

		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}

		m.User = User{
			ID:    msg.User.ID,
			Name:  msg.User.Name,
			Email: msg.User.Email,
		}

		byteMsg, err := json.Marshal(m)

		if err != nil {
			fmt.Println("marshal:", err)
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		fmt.Printf("recv: %s", byteMsg)

		err = channel.WriteMessage(messageType, byteMsg)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}
