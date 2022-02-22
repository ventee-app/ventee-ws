package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"

	"ventee-backend/configuration"
	"ventee-backend/types"
)

// store connections
var connections = []*types.ConnectionStruct{}

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
	registerConnectionData, _ := json.Marshal(types.RegisterConnectionDataStruct{
		ConnectionId: connectionId,
	})
	connection.WriteJSON(types.MessageStruct{
		Data:   string(registerConnectionData[:]),
		Event:  configuration.EVENTS.RegisterConnection,
		Issuer: configuration.BACKEND_ID,
		Target: connectionId,
	})

	// Store connection
	connectionStruct := new(types.ConnectionStruct)
	connectionStruct.Connection = connection
	connectionStruct.ConnectionId = connectionId
	connections = append(connections, connectionStruct)

	log.Println("connections size", len(connections))

	for {
		var parsedMessage types.MessageStruct
		parsingError := connection.ReadJSON(&parsedMessage)
		if parsingError != nil {
			errorMessageData, _ := json.Marshal(types.InvalidIncomingMessageStruct{
				Message: configuration.ERRORS.InvalidIncomingMessage,
			})
			connection.WriteJSON(types.MessageStruct{
				Data:   string(errorMessageData[:]),
				Event:  configuration.EVENTS.Error,
				Issuer: configuration.BACKEND_ID,
				Target: connectionId,
			})
		}

		if parsedMessage.Event == configuration.EVENTS.TransferContacts {
			log.Println(parsedMessage.Data)
			// handlers.TransferContacts(connection, connectionId, connections, parsedMessage)
		}
	}
}
