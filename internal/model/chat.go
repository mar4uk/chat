package model

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type Chat struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

var chats = []*Chat{
	{ID: 1, Title: "Default chat"},
}

func GetChat(id string) (*Chat, error) {
	chatID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.New("chat id is incorrect")
	}

	for _, c := range chats {
		if c.ID == chatID {
			return c, nil
		}
	}

	return nil, errors.New("chat not found")
}

func (c *Chat) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetChatMessages(c *Chat) []render.Renderer {
	list := []render.Renderer{}
	for _, m := range messages {
		if c.ID == m.ChatID {
			list = append(list, m)
		}
	}

	return list
}
