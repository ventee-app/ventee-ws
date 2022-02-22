package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"

	"ventee-backend/types"
)

func TransferContacts(
	connection *websocket.Conn,
	connectionId string,
	connections []*types.ConnectionStruct,
	message types.MessageStruct,
) {
	fmt.Println("is HERE", connection, connectionId, connections, message)
}
