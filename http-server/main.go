package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"golang-learning-project/http-server/internal/handler"
	"golang-learning-project/http-server/internal/service"
)

func main() {
	//propertiesConfig, err := config.NewConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}
	router := chi.NewRouter()
	// initialize service and handler
	// userRepo := repository.NewRepository(propertiesConfig.Username, propertiesConfig.Password)
	//userService := service.NewService(userRepo, propertiesConfig.ExternalApiToken)
	userService := service.NewService("propertiesConfig.ExternalApiToken")
	userHandler := handler.NewHandler(userService)

	router.MethodFunc(http.MethodPost, "/user", userHandler.Post)

	// decorate handler with logging functionality
	router.Route("/", func(subRouter chi.Router) {
		subRouter.Use(middleware.Logger)
		subRouter.MethodFunc(http.MethodGet, "/user/{id}", userHandler.GetUserById)
		subRouter.MethodFunc(http.MethodGet, "/users", userHandler.Get)
		subRouter.MethodFunc(http.MethodGet, "/generateNames", userHandler.GenerateNames)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Try to start server...")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdown)

	go func() {
		log.Println("Server started working...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-shutdown

	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped gracefully")
}
