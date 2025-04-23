package application

import (
	"github.com/gmalheiro/golang-playground-txdb/internal/http/handler"
	"github.com/gmalheiro/golang-playground-txdb/internal/repository"
	"github.com/gmalheiro/golang-playground-txdb/internal/route"
	"github.com/gmalheiro/golang-playground-txdb/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gmalheiro/golang-playground-txdb/database"
	"github.com/gmalheiro/golang-playground-txdb/internal/configs"
)

func Bootstrap() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("failed to load config file")
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("error while connecting to db, error: %s", err.Error())
	}

	db := database.GetConnection()

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.Heartbeat("/ping"))

	route.RegisterRoutes(r, userHandler)

	server := &http.Server{
		Addr:              ":" + configs.Get().ServerPort,
		Handler:           r,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Starting server at port %s", configs.Get().ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
