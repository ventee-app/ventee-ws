package types

import "github.com/gorilla/websocket"

type ConnectionStruct struct {
	ConnectionId string
	Connection   *websocket.Conn
}

type InvalidIncomingMessageStruct struct {
	Message string `json:"message"`
}

type MessageStruct struct {
	Data   string `json:"data"`
	Event  string `json:"event"`
	Issuer string `json:"issuer"`
	Target string `json:"target"`
}

type RegisterConnectionDataStruct struct {
	ConnectionId string `json:"connectionId"`
}

type TransferContacts struct {
	Contacts []interface{} `json:"contacts"`
	Target   string        `json:"target"`
}
