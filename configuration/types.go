package configuration

type EnvironmentsStruct struct {
	Development string
	Heroku      string
	Production  string
}

type EventsStruct struct {
	RegisterConnection string
	RequestContacts    string
	TransferComplete   string
	TransferContacts   string
}
