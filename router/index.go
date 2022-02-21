package router

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"

	"ventee-backend/configuration"
)

// store connections
var connections = []*ConnectionStruct{}

// Upgrade connection & disable origin check
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(request *http.Request) bool { return true },
}

func HandleConnection(writer http.ResponseWriter, request *http.Request) {
	connection, connectionError := upgrader.Upgrade(writer, request, nil)
	if connectionError != nil {
		return
	}

	defer connection.Close()

	// Create a new ID for connection & send it to the client
	connectionId := cuid.New()
	connection.WriteJSON(MessageStruct{
		Data: RegisterConnectionDataStruct{
			ConnectionId: connectionId,
		},
		Event: configuration.EVENTS.RegisterConnection,
	})

	// Store connection
	connectionStruct := new(ConnectionStruct)
	connectionStruct.Connection = connection
	connectionStruct.ConnectionId = connectionId
	connections = append(connections, connectionStruct)

	log.Println("connections size", len(connections))

	for {
		messageType, message, _ := connection.ReadMessage()
		log.Println("recv:", string(message[:]), messageType, connectionId)

		writeError := connection.WriteMessage(messageType, message)
		if writeError != nil {
			log.Println("write:", writeError)
			break
		}
	}
}
