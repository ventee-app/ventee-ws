package handler

type MessageStruct struct {
	Data  interface{} `json:"data"`
	Event string      `json:"event"`
}

type RegisterConnectionDataStruct struct {
	ConnectionId string `json:"connectionId"`
}
