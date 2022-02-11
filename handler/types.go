package handler

import "github.com/gorilla/websocket"

type ConnectionStruct struct {
	ConnectionId string
	Connection   *websocket.Conn
}

type MessageStruct struct {
	Data  interface{} `json:"data"`
	Event string      `json:"event"`
}

type RegisterConnectionDataStruct struct {
	ConnectionId string `json:"connectionId"`
}
