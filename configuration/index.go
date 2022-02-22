package configuration

const BACKEND_ID string = "backend"

const DEFAULT_PORT string = "9099"

var ENVIRONMENTS = EnvironmentsStruct{
	Development: "development",
	Heroku:      "heroku",
	Production:  "prouction",
}

var EVENTS = EventsStruct{
	RegisterConnection: "register-connection",
	RequestContacts:    "request-contacts",
	TransferComplete:   "transfer-complete",
	TransferContacts:   "transfer-contacts",
}
