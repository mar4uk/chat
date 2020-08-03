package model

import (
	"math/rand"
	"net/http"
)

type Message struct {
	ID        int64  `json:"id"`
	ChatID    int64  `json:"chatId"`
	UserID    int64  `json:"userId"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

var messages = []*Message{
	{ID: 1, ChatID: 1, UserID: 1, Text: "Hello!", CreatedAt: "2020-01-02T15:04:05-0700"},
	{ID: 2, ChatID: 1, UserID: 1, Text: "Hey!", CreatedAt: "2020-01-02T15:04:06-0700"},
}

func (m *Message) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateMessage(message *Message) (int64, error) {
	message.ID = rand.Int63()
	messages = append(messages, message)
	return message.ID, nil
}
