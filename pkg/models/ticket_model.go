package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    TicketNumber     string             `json:"ticketNumber,omitempty" validate:"required"`
    Price   int          `json:"price,omitempty" validate:"required"`
}