package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Problem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    User               `bson:"user_id"`
	Title     string             `bson:"title"`
	Question  string             `bson:"question"`
	TypeID    primitive.ObjectID `bson:"type_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	// Version   int                `bson:"__v"`
}
