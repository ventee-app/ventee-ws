package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"

	"ventee-backend/configuration"
)

// Upgrade connection & disable origin check
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(request *http.Request) bool { return true },
}

func Handle(writer http.ResponseWriter, request *http.Request) {
	connection, connectionError := upgrader.Upgrade(writer, request, nil)
	if connectionError != nil {
		log.Fatal(connectionError) // TODO: this should not be fatal
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

	// TODO: there should be different handlers depending on the event type
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
