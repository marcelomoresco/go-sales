package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Password string             `bson:"password,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Name     string             `bson:"name,omitempty"`
}