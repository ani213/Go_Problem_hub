package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProblemType struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string             `bson:"title" json:"title" validate:"required"`
	// Value     string             `bson:"value" json:"value" validate:"required"`
	Picture   string    `bson:"picture,omitempty" json:"picture,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type ProblemTypeInput struct {
	Title string `form:"title" json:"title" binding:"required"`
}
