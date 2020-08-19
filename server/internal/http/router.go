package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/auth"
	"github.com/mar4uk/chat/internal/ctxutils"
)

func middlewares() chi.Middlewares {
	return chi.Middlewares{
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
		render.SetContentType(render.ContentTypeJSON),
	}
}

type middleware func(next http.Handler) http.Handler

func chatMiddleware(a app.App) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			chatID, err := strconv.ParseInt(chi.URLParam(r, "chatID"), 10, 64)
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
			ctx := r.Context()
			chat, err := a.GetChat(ctx, chatID)
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}

			ctx = ctxutils.SetChat(ctx, chat)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func verifyJwtMiddleware(a app.App) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var header = r.Header.Get("Authorization")
			header = strings.Split(header, " ")[1]

			if header == "" {
				render.Render(w, r, &ErrResponse{HTTPStatusCode: 401, StatusText: "Unauthorized."})
				return
			}

			tk := &Token{}

			_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil {
				render.Render(w, r, ErrInternalServer(err))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func setupRouter(a app.App, ah auth.Auth) http.Handler {
	r := chi.NewRouter().With(middlewares()...)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Method(http.MethodPost, "/register", &registerUserHandler{auth: ah})
	r.Method(http.MethodPost, "/login", &loginUserHandler{auth: ah})

	r.Route("/chat/{chatID}", func(r chi.Router) {
		r.Use(verifyJwtMiddleware(a))
		r.Use(chatMiddleware(a))

		r.Method(http.MethodGet, "/messages", &getMessagesHandler{app: a})
		r.Method(http.MethodPost, "/messages", &createMessageHandler{app: a})
	})

	r.Method(http.MethodGet, "/socket", &websocketHandler{app: a})

	return r
}
