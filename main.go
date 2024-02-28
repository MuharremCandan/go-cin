package main

import (
	database "go-cin/db"
	"go-cin/router"
	"go-cin/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	pgDb, err := database.NewPgDb(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	pgDb.ConnectDb()

	gin.ForceConsoleColor()
	r := router.SetUpRouter()

	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal("Create Server Error: ", s.ListenAndServe())
}
