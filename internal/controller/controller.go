package controller

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chatController    chat
	messageController message
)

func Startup(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	chatController.registerRoutes(r)
	messageController.registerRoutes(r)
}
