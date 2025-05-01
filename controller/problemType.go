package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/ani213/Problemhub_backend/common"
	"github.com/ani213/Problemhub_backend/config"
	"github.com/ani213/Problemhub_backend/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProblemType(c *gin.Context) {
	var body model.ProblemTypeInput
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error-body": err.Error()})
		return
	}
	validate := common.ValidateSchema(c, body)
	if !validate {
		return
	}
	var problemTypesCollection = config.GetCollection("problem_types")
	newProblemTypes := model.ProblemType{
		ID:        primitive.NewObjectID(),
		Title:     body.Title,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	problemType, err := problemTypesCollection.InsertOne(ctx, newProblemTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"problemType": problemType,
	})
}
