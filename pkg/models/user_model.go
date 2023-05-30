package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
    Name     string             `json:"name,omitempty" validate:"required"`
    JobTitle    string             `json:"title,omitempty" validate:"required"`
		Email string `json:"email,omitempty" validate:"required"`
		Password string `json:"password,omitempty" validate:"required"`
		ConfirmPassword string `json:"confirmPassword,omitempty" validate:"required"`
		Tickets []primitive.ObjectID `json:"tickets,omitempty" validate:"required"`
		DepartmentId   primitive.ObjectID `json:"departmentId,omitempty" validate:"required"`
}