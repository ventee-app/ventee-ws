package configuration

type EnvironmentsStruct struct {
	Development string
	Heroku      string
	Prodduction string
}

type EventsStruct struct {
	RegisterConnection string
	RequestContacts    string
	TransferComplete   string
	TransferContacts   string
}
