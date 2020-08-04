package model

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
)

type Chat struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func GetChat(db *DB, id string) (*Chat, error) {
	chatID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.New("chat id is incorrect")
	}

	var chat *Chat

	if err = db.Collection("chats").FindOne(context.TODO(), bson.M{"id": chatID}).Decode(&chat); err != nil {
		log.Fatal(err)
		return nil, errors.New("chat not found")
	}

	return chat, nil
}

func (c *Chat) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetChatMessages(db *DB, c *Chat) []render.Renderer {
	list := []render.Renderer{}
	cursor, err := db.Collection("messages").Find(context.TODO(), bson.M{"chatId": c.ID})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var message Message
		if err = cursor.Decode(&message); err != nil {
			log.Fatal(err)
		}
		list = append(list, &message)
	}

	return list
}
