package controller

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mar4uk/chat/internal/model"
)

type chat struct{}

func (c chat) registerRoutes(r chi.Router) {
	r.Route("/chat", func(r chi.Router) {
		r.Route("/{chatID}", func(r chi.Router) {
			r.Use(chatContext)
			r.Get("/messages", getChatMessages)
		})
	})
}

func getChatMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chat, ok := ctx.Value("chat").(*model.Chat)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	messages := model.GetChatMessages(chat)
	render.Status(r, http.StatusOK)
	render.RenderList(w, r, messages)
}

func chatContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var chat *model.Chat
		var err error

		if chatID := chi.URLParam(r, "chatID"); chatID != "" {
			chat, err = model.GetChat(chatID)
		} else {
			render.Render(w, r, model.ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, model.ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "chat", chat)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
