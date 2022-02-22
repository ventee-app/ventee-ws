package configuration

type EnvironmentsStruct struct {
	Development string
	Heroku      string
	Prodduction string
}

type ErrorsStruct struct {
	InvalidIncomingMessage string
}

type EventsStruct struct {
	Error              string
	RegisterConnection string
	RequestContacts    string
	TransferComplete   string
	TransferContacts   string
}
