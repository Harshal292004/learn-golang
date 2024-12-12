package main

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	//  creates a URL with the web socket scheme (ws) pointing to a Web Socket server running on 8080 and using /ws path
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws",
	}

	// used to configure web socket connection params and the buffer size and writing data
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Dial() func attempts to establish ws conn to the server
	// returns 3 value conn , response err
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// sending the message   client sends a message "Hello server to the server with the write message function "
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// client waits for a response from the server using conn.ReadMessage()
	//returns 3 values
	//messageType (eg. binary or text)
	//message (actual msg sent by the server )
	//err

	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received:", message)

}
