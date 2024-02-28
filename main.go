package main

import (
	database "go-cin/db"
	"go-cin/handlers"
	"go-cin/repository"
	"go-cin/router"
	"go-cin/service"
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
		log.Fatalf("Failed to create databse connection: %v", err)
	}
	db, err := pgDb.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	gin.ForceConsoleColor()

	albumRepo := repository.NewAlbumRepository(db)
	albumService := service.NewAlbumService(albumRepo)
	albumHandler := handlers.NewAlbumHandler(albumService)
	r := router.SetUpRouter(albumHandler)

	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal("Create Server Error: ", s.ListenAndServe())
}
