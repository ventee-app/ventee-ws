package configuration

const DEFAULT_PORT = "9099"

var ENVIRONMENTS = EnvironmentsStruct{
	Development: "development",
	Heroku:      "heroku",
	Prodduction: "prouction",
}

var EVENTS = EventsStruct{
	RegisterConnection: "register-connection", // Backend -> Client
	RequestContacts:    "request-contacts",    // Client (receiver) -> Backend -> Client (sender)
	TransferComplete:   "transfer-complete",   // Client (receiver) -> Backend -> Client (sender)
	TransferContacts:   "transfer-contacts",   // Client (sender) -> Backend -> Client (receiver)
}
