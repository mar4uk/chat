package store

import (
	"context"

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
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var message Message
		if err = cursor.Decode(&message); err != nil {
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

func (m *mongoDatabase) GetMessageByID(ctx context.Context, id primitive.ObjectID) (Message, error) {
	var message Message

	if err := m.db.Collection("messages").FindOne(ctx, bson.M{"_id": id}).Decode(&message); err != nil {
		return Message{}, err
	}

	return message, nil
}

func (m *mongoDatabase) CreateMessage(ctx context.Context, message Message) (primitive.ObjectID, error) {
	message.ID = primitive.NewObjectID()
	if _, err := m.db.Collection("messages").InsertOne(ctx, message); err != nil {
		return primitive.NilObjectID, err
	}

	return message.ID, nil
}

func (m *mongoDatabase) CreateUser(ctx context.Context, user User) (primitive.ObjectID, error) {
	insertedResult, err := m.db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return insertedResult.InsertedID.(primitive.ObjectID), nil
}

func (m *mongoDatabase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user *User

	if err := m.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mongoDatabase) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	var user *User

	if err := m.db.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
