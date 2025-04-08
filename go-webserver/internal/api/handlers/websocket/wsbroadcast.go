package handler

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // map hold all active clients
var mutex = &sync.Mutex{}
var broadcast = make(chan []byte)

/*
WebSocket with broadcast to all connected clients
*/
func WsBroadcastHandler(c *gin.Context) {
	conn, err := upgraderBroadcast.Upgrade(c.Writer, c.Request, nil) // Upgrade the https request
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		http.Error(c.Writer, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	// closing connection when done
	defer conn.Close()

	// we will add this new clint to map
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	log.Printf("Client connected from %s", conn.RemoteAddr())

	// Setting a ping pong mechanism is used to keep the connection alive
	conn.SetPingHandler(func(string) error {
		conn.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(time.Second))
		return nil
	})

	// infinite loop for handling incoming messages
	for {
		// Read messages from connection
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			mutex.Lock()
			delete(clients, conn) // remove the client from map
			mutex.Unlock()
			break
		}

		log.Printf("Received message: %s", string(message))
		broadcast <- message // send the message to broadcast channel
	}
}

/*
This function will handle the broadcast of messages to all connected clients
It will run in a seprate go routine
*/
func handleBroadcast() {
	for {
		// Grad the lastes message from the channel
		message := <-broadcast
		mutex.Lock()
		// Loop through all clients and send the message
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error sending message to client: %v", err)
				client.Close()          // close the connection if error occurs
				delete(clients, client) // remove the client from map
			}
		}
		mutex.Unlock()
	}
}

// Upgrader is used to upgrade HTTP connection requests to WebSocket connections.
var upgraderBroadcast = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // allow all origins to connect
		return true
	},
}

// init function will be automatically called when the package is initialized
// and it will start the broadcast handler in a separate goroutine
func init() {
	go handleBroadcast()
	log.Println("WebSocket broadcast handler started")
}
