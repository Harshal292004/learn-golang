package CustomWebSocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for development
		return true
	},
}

// The websocket.Conn object represents an open WebSocket connection.
// WebSocket connections are stateful and can potentially hold a lot of data (buffers, connection state, etc.),
// which is why it's more efficient to return a pointer to the websocket.Conn rather than a copy of it.
// state modification ,since ws can have their state changed (sending or receiving data)
// we don't pass the Response Writer pointer as RW is a interface and in go interfaces are refrenced as pointers
// where as the Request is a struct object thus needs to be refrenced using pointer
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	// The line `Upgrader.CheckOrigin = func(r *http.Request) bool { return true }` is setting a custom
	// function for checking the origin of the WebSocket connection request. In this case, the function
	// always returns `true`, which means that any origin is allowed to establish a WebSocket connection
	// with the server.
	log.Println("We reached the log file here !!!!")
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket connection error:- ", err)
		return nil, err
	}

	return conn, nil
}
