package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mar4uk/chat/internal/controller"
)

func main() {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	controller.Startup(r)
	http.ListenAndServe(":8080", r)
}
