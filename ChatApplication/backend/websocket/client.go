package CustomWebSocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Represents a client connected to the WebSocket server.
type Client struct {
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex // used for synchronization
}

// A simple structure that represents the message being transmitted between clients
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		m := Message{Type: msgType, Body: string(msg)}

		c.Pool.BroadCast <- m
		log.Println("Msg Received===>>>\n", m)
		fmt.Println("msg received===>>>\n", m)
	}
}
