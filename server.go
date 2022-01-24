package main

import (
	"orchestra/actions"
	"orchestra/config"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{}

func main() {
	config.CreateServerEnvFile()
	// Get env variables
	_, port := getEnvVariables()

	http.HandleFunc("/", webSocketHandler)
	log.Println("Paste -> wss://localhost" + port + "/ <- in the respective field")
	log.Fatal(http.ListenAndServeTLS(port, "./certificates/localhost.pem", "./certificates/localhost-key.pem", nil))
}

func webSocketHandler(response http.ResponseWriter, request *http.Request) {
	// Handle cors
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// Upgrade all the connections to websocket connections
	ws, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println("An error occurred:", err)
	}
	log.Println("Successfully connected!")
	// Start listening and reading all the incoming connections
	connectionsReader(ws)
}

func connectionsReader(conn *websocket.Conn) error {
	// The struct helps to parse and access
	// the data received from client
	type Message struct {
		Event string
		PosX  float64
		PosY  float64
		Key   string
	}
	var (
		receivedMessage string
		message         Message
	)
	// Send a message to client
	msg := []byte("Message from the server")
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Fatal("An error occurred:", err)
		return err
	}
	// Keep listening for connections
	for {
		// Read each connection
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("An error occurred ->", err)
			return err
		}

		// Get the client data
		// received in stringified JSON
		receivedMessage = string(p)
		// turn the stringified JSON into a Message struct
		// so we can access the data more easily
		json.Unmarshal([]byte(receivedMessage), &message)
		fmt.Println("coordinates", message.PosX, message.PosY)
		switch message.Event {
		case "mousemove":
			actions.MouseMove(message.PosX, message.PosY)
			break
		case "mouse-click":
			actions.MouseClick()
			break
		case "type":
			actions.KeyType(message.Key)
			break
		case "scroll":
			actions.Scroll(message.PosX, message.PosY)
			break
		}
	}
}

// Read the generated env file
func getEnvVariables() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("An error occurred loading the environment variables")
	}
	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	return host, port
}
