package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ani213/Problemhub_backend/common"
	"github.com/ani213/Problemhub_backend/config"
	"github.com/ani213/Problemhub_backend/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    string
	Role  string
	Email string
}

func Register(c *gin.Context) {
	var body model.RegisterBody
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := c.ShouldBindBodyWithJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	var userCollection = config.GetCollection("users")

	err = userCollection.FindOne(ctx, bson.M{"$or": []bson.M{
		{"email": body.Email},
		{"username": body.Username},
	}}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		salt, err := common.GenerateSalt()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashPassword := common.EncryptPassword(body.Password, salt)
		password := model.Password{
			Salt:             salt,
			HashedPassword:   hashPassword,
			VerificationCode: 000000,
		}
		newUser := model.User{
			ID:        primitive.NewObjectID(),
			Username:  body.Username,
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Password:  password,
			Role:      "user",
		}
		user, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username or email exiest"})
		return
	}

}

func Users(c *gin.Context) {
	var userCollection = config.GetCollection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cusers, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cusers.Close(ctx)
	var users []model.User
	if err = cusers.All(ctx, &users); err != nil {
		log.Println("Error decoding users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing user data"})
		return
	}

	// Return the list of users as JSON
	c.JSON(http.StatusOK, users)

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var userCollection = config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	err := userCollection.FindOne(ctx, bson.M{"email": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username or password is wrong"})
		return
	}
	// defer cusers.Close(c)
	isCheckedPassword := common.CheckPassword(password, user.Password.HashedPassword, user.Password.Salt)
	if !isCheckedPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username or password is wrong"})
		return
	}
	tokenData := map[string]interface{}{
		"id":       user.ID.Hex(),
		"role":     user.Role,
		"username": user.Username,
		"email":    user.Email,
	}

	accessToken, err := common.GenerateAccessToken(tokenData)
	if err != nil {
		model.SendError(c, model.SomethingWentWrong(err))
		return
	}
	refreshToken, err := common.GenerateRefreshToken(tokenData)
	if err != nil {
		model.SendError(c, model.SomethingWentWrong(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"expire_refresh": 1800, "expire_access": 2400, "access_token": accessToken, "refresh_token": refreshToken})
}
