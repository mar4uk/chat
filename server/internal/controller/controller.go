package controller

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chatController    chat
	messageController message
)

func Startup(router chi.Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	chatController.registerRoutes(router)
	messageController.registerRoutes(router)
}
