package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    TicketNumber     string             `json:"ticketNumber,omitempty" validate:"required"`
		UserID   primitive.ObjectID `json:"userId,omitempty" validate:"required"`
}