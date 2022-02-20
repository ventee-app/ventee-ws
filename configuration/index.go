package configuration

const DEFAULT_PORT = "9099"

var ENVIRONMENTS = EnvironmentsStruct{
	Development: "development",
	Heroku:      "heroku",
	Prodduction: "prouction",
}

var EVENTS = EventsStruct{
	RegisterConnection: "register-connection",
}
