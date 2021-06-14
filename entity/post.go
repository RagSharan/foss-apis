package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	PostId  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId  string             `json:"userId" bson:"userId"`
	Content string             `json:"content" bson:"content"`
	Comment Comment            `json:"comment" bson:"comment"`
	Vote    Vote               `json:"vote" bson:"vote"`
}
