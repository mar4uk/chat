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
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

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
