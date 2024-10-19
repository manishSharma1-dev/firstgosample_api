package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id        primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	MovieName string             `bson:"moviename,omitempty" json:"moviename,omitempty"`
	Watched   bool               `bson:"watched,omitempty" json:"watched,omitempty"`
}
