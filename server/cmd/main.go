package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mar4uk/chat/internal/controller"
	"github.com/mar4uk/chat/internal/model"
)

func main() {
	client, err := model.NewClient("mongodb://root:password@localhost:27017")

	if err != nil {
		log.Fatal(err)
	}
	db := &model.DB{client.Database("chat")}
	defer client.Disconnect(context.TODO())

	fmt.Println("Connected to MongoDB!")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	controller.Startup(r)
	http.ListenAndServe(":8080", r)
}
