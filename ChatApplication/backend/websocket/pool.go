package CustomWebSocket

import (
	"fmt"
	"log"
)

// Pool is a container for all ws clients
// maintains a map where client is the key and a bool value is used to represent whether the client is active
// It has channels
// Register : For new clients joining the pool
// Unregister: For clients disconnecting
// BroadCast: For broadcasting messages to all clients
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	BroadCast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		BroadCast:  make(chan Message),
	}

}

// this is a goroutine that continously process the incomming messages on these cahnnels,handling
// Registerring new clients and sending a welcome message to all otheres
// UnRegestering cleints and notiftinf others
// Broadcasting messages to all clients when a new message is received
func (pool *Pool) Start() {
	log.Println("Started the pooling for the routine")
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("total connection pool:- ", len(pool.Clients))
			for k, _ := range pool.Clients {
				fmt.Println(k)
				log.Println("I actually do register users")
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined"})
			}

		case client := <-pool.Unregister:
			//map => key
			delete(pool.Clients, client)
			fmt.Println("total connection pool:- ", len(pool.Clients))
			for k, _ := range pool.Clients {
				fmt.Println(k)
				log.Println("unregistered ,", k)
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected"})
			}

		case msg := <-pool.BroadCast:
			fmt.Println("broadcasting a message")
			log.Println("Broadcasting a message")
			for k, _ := range pool.Clients {
				if err := k.Conn.WriteJSON(msg); err != nil {
					fmt.Println(err)
					return
				} else {
					log.Println("msg is ", msg)
				}
			}
		}
	}
}
