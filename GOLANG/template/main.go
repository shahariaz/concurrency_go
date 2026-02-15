package main

import (
	"html/template"
)

type Email struct {
	To      string
	From    string
	Subject string
	Body    string
	Items   []string
	Unread  int
}

type SWriter struct{}

func (s *SWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	err = nil
	return
}
func main() {

	templ := `
	   To: {{.To}}
	   From: {{.From}}
	   Subject: {{.Subject}}
	   Body: {{.Body}}
	   Items:
	   {{range .Items}}
	       - {{.}}
		{{end}}
	{{if gt .Unread 0}}
	   You have {{.Unread}} unread messages.
	{{else}}
	   You have no unread messages.
	{{end}}
	`
	temp, err := template.New("email-message").Parse(templ)
	if err != nil {
		panic(err)
	}

	email := Email{
		To:      "user@example.com",
		From:    "admin@example.com",
		Subject: "Important Message",
		Body:    "This is an important message.",
		Items:   []string{"Item 1", "Item 2", "Item 3"},
		Unread:  5,
	}

	err = temp.Execute(&SWriter{}, email)
	if err != nil {
		panic(err)
	}
}
