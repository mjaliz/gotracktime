package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/handlers"
	"github.com/mjaliz/gotracktime/internal/middlewares"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.Repo.Home)
	r.POST("/user/signUp", handlers.Repo.SignUp)
	r.POST("/user/signIn", handlers.Repo.SignIn)
	r.Use(middlewares.Auth())
	r.GET("/ping", handlers.Ping)
	r.POST("/timeEntity", handlers.Repo.CreateTimeEntity)
	return r
}
