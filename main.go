package main

import (
	"go-cin/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	r := router.SetUpRouter()

	s := &http.Server{
		Addr:           ":8089",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal("Create Server Error: ", s.ListenAndServe())
}
