package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mar4uk/chat/internal/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type message struct{}

func (m message) registerRoutes(r chi.Router) {
	r.Route("/messages", func(r chi.Router) {
		r.Post("/", createMessage)
	})
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	var m *model.Message

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, ok := r.Context().Value("db").(*model.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	model.CreateMessage(db, m)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, m)
}
