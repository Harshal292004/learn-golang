package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	//web socket upgrader
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to a WebSocket connection.
		//note that a websocket is a state full protocol
		//where as the http is a state less protocol
		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Received:", message)

			err = conn.WriteMessage(messageType, []byte("Hello,Client !"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})
	// Start the server
	http.ListenAndServe(":8080", nil)
}
