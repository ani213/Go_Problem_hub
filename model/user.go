package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Role      string             `bson:"role" json:"role"`
	Password  Password           `bson:"password" json:"password"`
	Status    string             `bson:"status" json:"status"`
	Picture   string             `bson:"picture" json:"picture"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	// Version   int                `bson:"__v" json:"version"`
}

// Password represents the password structure
type Password struct {
	Salt             string `bson:"salt" json:"salt"`
	HashedPassword   string `bson:"hashedPassword" json:"hashedPassword"`
	VerificationCode int    `bson:"varificationCode" json:"varificationCode"`
}

type RegisterBody struct {
	Username  string `json:"username" validate:"required,min=3,max=20"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}
