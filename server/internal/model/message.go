package model

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChatID    int64              `json:"chatId" bson:"chatId"`
	UserID    int64              `json:"userId" bson:"userId"`
	Text      string             `json:"text" bson:"text"`
	CreatedAt string             `json:"createdAt" bson:"createdAt"`
}

func (m *Message) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateMessage(db *DB, message *Message) (primitive.ObjectID, error) {
	message.ID = primitive.NewObjectID()
	insertResult, err := db.Collection("messages").InsertOne(context.TODO(), &message)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID.(primitive.ObjectID), nil
}
