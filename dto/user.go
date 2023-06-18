package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDto struct {
	ID        primitive.ObjectID `json:"-"`
	Username  string             `json:"Username" binding:"required"`
	CreatedAt string             `json:"CreatedAt"`
	UpdatedAt string             `json:"UpdatedAt"`
}
