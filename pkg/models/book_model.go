package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Tickets []primitive.ObjectID `json:"tickets,omitempty" validate:"required"`
    UserID   primitive.ObjectID `json:"userId,omitempty" validate:"required"`
    StartDate time.Time          `json:"startDate,omitempty" validate:"required"`
	EndDate   time.Time          `json:"endDate,omitempty" validate:"required"`
}