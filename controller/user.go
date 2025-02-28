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
)

type User struct {
	ID    string
	Role  string
	Email string
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get request",
	})
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
	}
	// defer cusers.Close(c)
	isCheckedPassword := common.CheckPasswor(password, user.Password.HashedPassword, user.Password.Salt)
	if !isCheckedPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username or password is wrong"})
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
	}
	refreshToken, err := common.GenerateRefreshToken(tokenData)
	if err != nil {
		model.SendError(c, model.SomethingWentWrong(err))
	}
	c.JSON(http.StatusOK, gin.H{"expire_refresh": 1800, "expire_access": 2400, "access_token": accessToken, "refresh_token": refreshToken})
}
