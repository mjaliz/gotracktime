package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjaliz/gotracktime/internal/handlers"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.Repo.Home)
	r.POST("/user/sign_up", handlers.Repo.SignUp)
	return r
}
