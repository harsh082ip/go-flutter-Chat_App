package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `json:"name"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	Username        string             `json:"username"`
	UserId          string             `json:"userid"`
	Profile_Pic_Url string             `json:"profile_pic_url"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RecentlyViewed struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserIDs []string           `json:"userIDs"`
}

type RecentChats struct {
	ID        primitive.ObjectID `bson:"_id"`
	Usernames []string           `json:"usernames"`
}
