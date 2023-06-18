package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Username  string             `json:"username" validate:"required"`
	Password  string             `json:"password" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
