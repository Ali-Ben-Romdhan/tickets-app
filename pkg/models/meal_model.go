package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meal struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    TicketReference     string             `json:"ticketReference,omitempty" validate:"required"`
    UserId string             `json:"userId,omitempty" validate:"required"`
    Title    string             `json:"title,omitempty" validate:"required"`
}