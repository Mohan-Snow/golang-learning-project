package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/getNames", getNames)

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

func getNames(writer http.ResponseWriter, request *http.Request) {
	names := map[int]string{0: "Mohan", 1: "Dinos", 2: "Some Name"}

	for key, val := range names {
		temp := fmt.Sprintf("Id=%d Name=%s", key, val)
		fmt.Println(temp)
		// TODO: Handle error
		fmt.Fprintln(writer, temp)
	}
}
