package auth

import (
	"context"

	"github.com/mar4uk/chat/internal/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auth is an interface for auth
type Auth interface {
	CreateUser(ctx context.Context, user User) (primitive.ObjectID, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type auth struct {
	db store.Database
}

// User is
type User struct {
	ID       primitive.ObjectID
	Name     string
	Email    string
	Password string
}

// NewAuth is chat initialization function
func NewAuth(db store.Database) Auth {
	return &auth{
		db: db,
	}
}

func (a *auth) CreateUser(ctx context.Context, user User) (primitive.ObjectID, error) {
	userID, err := a.db.CreateUser(ctx, store.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return userID, err
	}

	return userID, nil
}

func (a *auth) GetUserByEmail(ctx context.Context, email string) (User, error) {
	dbUser, err := a.db.GetUserByEmail(ctx, email)
	user := User{}

	if err != nil {
		return user, err
	}

	return User{
		ID:       dbUser.ID,
		Name:     dbUser.Name,
		Email:    dbUser.Email,
		Password: dbUser.Password,
	}, nil
}
