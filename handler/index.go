package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"

	"ventee-ws/configuration"
)

var connections []*ConnectionStruct

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(request *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnection(writer http.ResponseWriter, request *http.Request) {
	connection, connectionError := upgrader.Upgrade(writer, request, nil)
	if connectionError != nil {
		return
	}

	defer connection.Close()

	connectionId := cuid.New()
	registerConnectionData, _ := json.Marshal(RegisterConnectionDataStruct{
		ConnectionId: connectionId,
	})
	connection.WriteJSON(MessageStruct{
		Data:   string(registerConnectionData[:]),
		Event:  configuration.EVENTS.RegisterConnection,
		Issuer: configuration.BACKEND_ID,
		Target: connectionId,
	})

	connectionStruct := new(ConnectionStruct)
	connectionStruct.Connection = connection
	connectionStruct.ConnectionId = connectionId
	connections = append(connections, connectionStruct)

	log.Println("Connected", connectionId, "| Total:", len(connections))

	for {
		var parsedMessage MessageStruct
		parsingError := connection.ReadJSON(&parsedMessage)
		if parsingError != nil {
			var index int = -1
			for i := range connections {
				if connections[i].ConnectionId == connectionId {
					index = i
				}
			}
			if index >= 0 {
				connections[index] = connections[len(connections)-1]
				connections = connections[:len(connections)-1]
				log.Println("Disconnected", connectionId, "| Total:", len(connections))
			}
			break
		}

		if parsedMessage.Event == configuration.EVENTS.RequestContacts &&
			parsedMessage.Issuer != "" && parsedMessage.Target != "" {
			var target *ConnectionStruct
			for i := range connections {
				if connections[i].ConnectionId == parsedMessage.Target {
					target = connections[i]
				}
			}
			if target != nil {
				target.Connection.WriteJSON(MessageStruct{
					Event:  configuration.EVENTS.RequestContacts,
					Issuer: connectionId,
					Target: target.ConnectionId,
				})
			}
		}

		if parsedMessage.Event == configuration.EVENTS.TransferContacts &&
			parsedMessage.Issuer != "" && parsedMessage.Target != "" &&
			parsedMessage.Data != "" {
			var target *ConnectionStruct
			for i := range connections {
				if connections[i].ConnectionId == parsedMessage.Target {
					target = connections[i]
				}
			}
			if target != nil {
				target.Connection.WriteJSON(MessageStruct{
					Data:   parsedMessage.Data,
					Event:  configuration.EVENTS.TransferContacts,
					Issuer: connectionId,
					Target: target.ConnectionId,
				})
			}
		}

		if parsedMessage.Event == configuration.EVENTS.TransferComplete &&
			parsedMessage.Issuer != "" && parsedMessage.Target != "" {
			var target *ConnectionStruct
			for i := range connections {
				if connections[i].ConnectionId == parsedMessage.Target {
					target = connections[i]
				}
			}
			if target != nil {
				target.Connection.WriteJSON(MessageStruct{
					Event:  configuration.EVENTS.TransferComplete,
					Issuer: connectionId,
					Target: target.ConnectionId,
				})
			}
		}
	}
}
