package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
		UserID   primitive.ObjectID `json:"userId" bson:"userId"`
		TicketID primitive.ObjectID `json:"ticketId" bson:"ticketId"`
		BookID primitive.ObjectID `json:"bookId" bson:"ticketId"`
		StartDate time.Time          `json:"startDate,omitempty" validate:"required"`
		EndDate   time.Time          `json:"endDate,omitempty" validate:"required"`
}