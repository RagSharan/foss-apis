package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	UserId    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Contact   string             `json:"contact" bson:"contact"`
	PostCount int                `json:"postcount" bson:"postcount"`
	TotalView int                `json:"totalview" bson:"totalview"`
}
