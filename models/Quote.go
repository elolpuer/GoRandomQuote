package models

//Quote ...
type Quote struct {
	ID     int
	Author string
	Body   string
}

//AddQuote ...
func AddQuote(id int, author string, body string) *Quote {
	return &Quote{id, author, body}
}
