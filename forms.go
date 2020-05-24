// forms.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

var details ContactDetails

func info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "_info_\n")
	fmt.Fprintf(w, "Email: %s\n", details.Email)
	fmt.Fprintf(w, "Subject: %s\n", details.Subject)
	fmt.Fprintf(w, "Message: %s\n", details.Message)
}

func reader(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details.Email = r.FormValue("email")
	details.Subject = r.FormValue("subject")
	details.Message = r.FormValue("message")

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {

	http.HandleFunc("/info", info)
	http.HandleFunc("/", reader)

	http.ListenAndServe(":4080", nil)
}
