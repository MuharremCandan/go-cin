package router

import (
	"go-cin/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	return r
}
