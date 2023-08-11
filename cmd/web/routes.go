package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/handlers"
	"github.com/mjaliz/gotracktime/internal/middlewares"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.Repo.Home)
	r.POST("/user/sign_up", handlers.Repo.SignUp)
	r.POST("/user/sign_in", handlers.Repo.SignIn)
	r.Use(middlewares.Auth())
	r.GET("/ping", handlers.Ping)
	return r
}
