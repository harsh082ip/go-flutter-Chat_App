package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBMessage struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

type Conversation struct {
	ID       primitive.ObjectID `bson:"_id"`
	Messages []DBMessage        `json:"messages"`
}
