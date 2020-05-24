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

func info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "_info_\n")
}

func reader(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {

	http.HandleFunc("/info", info)
	http.HandleFunc("/", reader)

	http.ListenAndServe(":4080", nil)
}
