package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Response struct {
	Greeting string `json:"greeting"`
	Time string `json:"time"`
	Ip string `json:"ip"`
}

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
	greeting := GetGreetingFromRequest(r) + " from backend"
	timeString := time.Now().String()  
	ip := GetRealIPFromRequest(r)

	response := Response{greeting, timeString, ip}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetGreetingFromRequest(r *http.Request) string {
	keys, ok := r.URL.Query()["greeting"]
	if ok {
        return keys[0]
    }
	return ""
}

func GetRealIPFromRequest(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-IP")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarder-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
	
    return strings.Split(IPAddress, ":")[0]
}