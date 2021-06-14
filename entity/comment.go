package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	UserId  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Postid  string             `json:"postId" bson:"postId"`
	Content string             `json:"content" bson:"content"`
}
