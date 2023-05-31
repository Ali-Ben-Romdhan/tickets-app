package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Tickets []primitive.ObjectID `json:"tickets,omitempty" validate:"required"`
    UserID   primitive.ObjectID `json:"userId,omitempty" validate:"required"`
    Price   int          `json:"price,omitempty" validate:"required"`
}