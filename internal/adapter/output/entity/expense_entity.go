package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseEntity struct {
	ID              primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Description     string               `bson:"description,omitempty"`
	Price     		float64              `bson:"price,omitempty"`
	CompanyName     string               `bson:"company_name,omitempty"`
	DateTime     	time.Time            `bson:"datetime,omitempty"`
}