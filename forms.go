// forms.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "_info_\n")
}

func main() {

	type ContactDetails struct {
		Email   string
		Subject string
		Message string
	}

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/info", info)

	http.ListenAndServe(":4080", nil)
}
