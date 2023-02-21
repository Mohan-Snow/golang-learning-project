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
)

func main() {
	router := chi.NewRouter()
	router.Get("/user/{id}", getUserById)
	router.Get("/getNames", getNames)
	userHandler := &UserHandler{data: make(map[int]string)}
	router.MethodFunc(http.MethodGet, "/users", userHandler.Get)
	router.MethodFunc(http.MethodPost, "/user", userHandler.Post)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	//http.HandleFunc("/getNames", getNames)
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
