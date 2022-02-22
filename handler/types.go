package handler

import "github.com/gorilla/websocket"

type ConnectionStruct struct {
	ConnectionId string
	Connection   *websocket.Conn
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
