package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
		Departments []primitive.ObjectID `json:"departments,omitempty" validate:"required"`
    Name   string          `json:"price,omitempty" validate:"required"`
}