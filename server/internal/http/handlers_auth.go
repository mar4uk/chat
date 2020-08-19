package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"github.com/mar4uk/chat/internal/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User struct declaration
type User struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
}

// Token struct declaration
type Token struct {
	UserID primitive.ObjectID
	Name   string
	Email  string
	*jwt.StandardClaims
}

// LoginResponse struct declaration
type LoginResponse struct {
	User  *User
	Token string `json:"token"`
}

type registerUserHandler struct {
	auth auth.Auth
}

type loginUserHandler struct {
	auth auth.Auth
}

var jwtSecret = "secret"

func (h *registerUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user *User

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	user.Password = string(pass)

	_, err = h.auth.CreateUser(ctx, auth.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, user)
}

func (h *loginUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u *User
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	user, err := h.auth.GetUserByEmail(ctx, u.Email)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		render.Render(w, r, &ErrResponse{
			Err:            err,
			HTTPStatusCode: 401,
			StatusText:     "Wrong email or password",
			ErrorText:      err.Error(),
		})
		return
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk := &Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, LoginResponse{
		User: &User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Token: tokenString,
	})
}
