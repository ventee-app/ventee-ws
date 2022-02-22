package configuration

const BACKEND_ID string = "backend"

const DEFAULT_PORT string = "9099"

var ENVIRONMENTS = EnvironmentsStruct{
	Development: "development",
	Heroku:      "heroku",
	Prodduction: "prouction",
}

var ERRORS = ErrorsStruct{
	InvalidIncomingMessage: "INVALID_INCOMING_MESSAGE",
}

var EVENTS = EventsStruct{
	Error:              "error",               // Backend -> Client
	RegisterConnection: "register-connection", // Backend -> Client
	RequestContacts:    "request-contacts",    // Client (receiver) -> Backend -> Client (sender)
	TransferComplete:   "transfer-complete",   // Client (receiver) -> Backend -> Client (sender)
	TransferContacts:   "transfer-contacts",   // Client (sender) -> Backend -> Client (receiver)
}
