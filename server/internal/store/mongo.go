package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoDatabase creates new instance of Database for MongoDB
func NewMongoDatabase(ctx context.Context, uri string, dbName string) (Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &mongoDatabase{
		client: client,
		db:     db,
	}, nil
}

func (m *mongoDatabase) Close(ctx context.Context) {
	m.client.Disconnect(ctx)
}

func (m *mongoDatabase) GetMessages(ctx context.Context, chatID int64) ([]Message, error) {
	var list []Message
	cursor, err := m.db.Collection("messages").Find(ctx, bson.M{"chatId": chatID})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var message Message
		if err = cursor.Decode(&message); err != nil {
			log.Fatal(err)
			return nil, err
		}
		list = append(list, message)
	}

	return list, nil
}

func (m *mongoDatabase) GetChat(ctx context.Context, chatID int64) (*Chat, error) {
	var chat *Chat

	if err := m.db.Collection("chats").FindOne(ctx, bson.M{"id": chatID}).Decode(&chat); err != nil {
		return nil, err
	}
	return chat, nil
}

func (m *mongoDatabase) CreateMessage(ctx context.Context, message Message) error {
	message.ID = primitive.NewObjectID()
	if _, err := m.db.Collection("messages").InsertOne(ctx, message); err != nil {
		return err
	}

	return nil
}
