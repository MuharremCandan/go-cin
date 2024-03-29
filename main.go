package main

import (
	"go-cin/cloud"
	database "go-cin/db"
	"go-cin/handlers"
	"go-cin/pkg/config"
	"go-cin/repository"
	"go-cin/router"
	"go-cin/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig("./")
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

	cld, err := cloud.NewCloud(config.Cloud.CloudinaryURL)
	if err != nil {
		log.Fatalf("Failed to create cloudinary connection: %v", err)
	}
	gin.ForceConsoleColor()

	r := gin.Default()

	albumRepo := repository.NewAlbumRepository(db)
	albumService := service.NewAlbumService(albumRepo)
	albumHandler := handlers.NewAlbumHandler(albumService, cld)

	router.NewRouter(albumHandler).SetUpRouter(r)

	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal("Create Server Error: ", s.ListenAndServe())
}
