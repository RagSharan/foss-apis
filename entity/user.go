package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	UserId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
}
