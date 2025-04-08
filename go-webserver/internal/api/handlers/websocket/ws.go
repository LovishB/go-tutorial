package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketHandler handles WebSocket connections.
func WsHandler(c *gin.Context) {
	// Upgrade the connection request to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		http.Error(c.Writer, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close() // all connections should be closed when done(defer)

	log.Printf("Client connected from %s", conn.RemoteAddr())

	// Setting a ping pong mechanism is used to keep the connection alive
	conn.SetPingHandler(func(string) error {
		conn.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(time.Second))
		return nil
	})

	// Since it's a continious connection, we use infinite loop to handle incoming messages
	for {
		// Read messages from connection
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break // If error occurs, break the loop
		}

		log.Printf("Received message: %s", string(message))

		// Echo the message back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break // If error occurs, break the loop
		}
	}
}

// Upgrader is used to upgrade HTTP connection requests to WebSocket connections.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // allow all origins to connect
		return true
	},
}
