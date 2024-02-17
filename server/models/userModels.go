package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `json:"name"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	Username        string             `json:"username"`
	UserId          string             `json:"userid"`
	Profile_Pic_Url string             `json:"profile_pic_urlllll"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RecentlyViewed struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserIDs []string           `json:"userIDs"`
}
