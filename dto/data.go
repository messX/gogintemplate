package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataDto struct {
	ID   primitive.ObjectID `json:"-"`
	Name string             `json:"Name"`
}
