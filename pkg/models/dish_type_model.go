package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DishType struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Dishes []primitive.ObjectID `json:"dishes,omitempty" validate:"required"`
    Name    string             `json:"title,omitempty" validate:"required"`
}