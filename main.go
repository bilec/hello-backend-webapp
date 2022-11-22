package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var port = "8080"

func init() {
	var envPort = os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}
}

func main() {
	log.Println("Starting hello-backend-webapp on port " + port + " ...")

	http.HandleFunc("/backend", backedHandler)

	server := http.Server{Addr: ":" + port}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	server.Shutdown(context.Background())
}

func backedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from backend"))
}