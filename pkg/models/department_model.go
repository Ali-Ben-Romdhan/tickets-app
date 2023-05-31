package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    CompanyID   primitive.ObjectID `json:"companyId,omitempty" validate:"required"`
    Name   string          `json:"price,omitempty" validate:"required"`
}