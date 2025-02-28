package routes

import (
	"github.com/ani213/Problemhub_backend/controller"
	"github.com/ani213/Problemhub_backend/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	public := router.Group("/")
	public.POST("/login", controller.Login)
	public.GET("/register", controller.Register)

	/***********************/
	protected := router.Group("/v1")
	protected.Use(middleware.Authentication())
	protected.GET("/user", controller.Users)
	protected.GET("/problems", controller.GetAllProblems)
}
