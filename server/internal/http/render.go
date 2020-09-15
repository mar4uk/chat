package http

import (
	"net/http"

	"github.com/go-chi/render"
)

// Render method for message
func (m *Message) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// MessagesListResponse is render method for array of messages
func MessagesListResponse(messages []*Message) []render.Renderer {
	list := []render.Renderer{}
	for _, message := range messages {
		list = append(list, message)
	}
	return list
}

// Render method for user
func (m *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render method for UserResponse
func (resp UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
