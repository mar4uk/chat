package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/auth"
	"github.com/mar4uk/chat/internal/ctxutils"
	"github.com/mar4uk/chat/internal/logger"
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

func chatMiddleware(a app.App) func(next http.Handler) http.Handler {
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

func getTokenStringFromRequest(r *http.Request) string {
	var tokenString = r.URL.Query().Get("jwt")

	var header = r.Header.Get("Authorization")

	if header != "" {
		tokenString = strings.Split(header, " ")[1]
	}

	return tokenString
}

func loggerMiddleware(logger *logger.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(logger)
}

func verifyJwtMiddleware(a app.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenString := getTokenStringFromRequest(r)

			if tokenString == "" {
				render.Render(w, r, &ErrResponse{HTTPStatusCode: 401, StatusText: "Unauthorized."})
				return
			}

			tk := &Token{}

			token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil {
				render.Render(w, r, ErrInternalServer(err))
				return
			}

			ctx := r.Context()

			if claims, ok := token.Claims.(*Token); ok && token.Valid {
				ctx = ctxutils.SetUser(ctx, &app.User{
					ID:    claims.UserID,
					Name:  claims.Name,
					Email: claims.Email,
				})
			} else {
				render.Render(w, r, ErrInternalServer(err))
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func setupRouter(a app.App, ah auth.Auth, logger *logger.Logger) http.Handler {
	r := chi.NewRouter().With(middlewares()...)
	r.Use(loggerMiddleware(logger))
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("this is panic!")
	})

	r.Method(http.MethodPost, "/register", &registerUserHandler{auth: ah})
	r.Method(http.MethodPost, "/login", &loginUserHandler{auth: ah})

	r.Options("/login", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})
	r.Options("/register", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})

	r.Route("/user", func(r chi.Router) {
		r.Use(verifyJwtMiddleware(a))
		r.Method(http.MethodGet, "/", &getUserHandler{auth: ah})
	})

	r.Route("/chat/{chatID}", func(r chi.Router) {
		r.Use(verifyJwtMiddleware(a))
		r.Use(chatMiddleware(a))

		r.Method(http.MethodGet, "/messages", &getMessagesHandler{app: a})
		r.Method(http.MethodPost, "/messages", &createMessageHandler{app: a})
	})

	r.With(verifyJwtMiddleware(a)).Method(http.MethodGet, "/socket", &websocketHandler{app: a})

	return r
}
