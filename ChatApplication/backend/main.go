package main

import (
	CustomWebSocket "chatapp/websocket"
	"log"
	"net/http"
	"os"
)

func serverWs(pool *CustomWebSocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket connection attempt received")
	log.Printf("Request details - Method: %s, Host: %s, Remote Addr: %s",
		r.Method, r.Host, r.RemoteAddr)

	conn, err := CustomWebSocket.Upgrade(w, r)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		http.Error(w, "Could not upgrade WebSocket", http.StatusBadRequest)
		return
	}

	log.Println("WebSocket connection successfully upgraded")

	client := &CustomWebSocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
}

func setupRoutes() {
	log.Println("Route Setup started...")
	pool := CustomWebSocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("WebSocket connection attempt")
		serverWs(pool, w, r)
	})
}

func main() {
	// Open a log file and create it if it doesn't exist
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Set the log output to the file
	log.SetOutput(logFile)

	setupRoutes()
	log.Println("Server started on ws://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
