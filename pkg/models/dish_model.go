package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dish struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    DishTypeId string             `json:"dishTypeId,omitempty" validate:"required"`
    Name    string             `json:"title,omitempty" validate:"required"`
}