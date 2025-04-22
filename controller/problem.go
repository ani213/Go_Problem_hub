package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/ani213/Problemhub_backend/config"
	"github.com/ani213/Problemhub_backend/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProblems(c *gin.Context) {
	var problemCollection = config.GetCollection("problems")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cproblem, err := problemCollection.Find(ctx, bson.M{})
	if err != nil {
		model.SendError(c, model.SomethingWentWrong(err))
	}
	defer cproblem.Close(ctx)
	var problems []model.Problem
	if err = cproblem.All(ctx, &problems); err != nil {
		model.SendError(c, model.SomethingWentWrong(err))
		return
	}
	c.JSON(http.StatusOK, problems)
}

func AddProblem(c *gin.Context) {

	// var body model.Problem
	// // var problemCollection = config.GetCollection("problems")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

}
