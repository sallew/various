package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name string
}

var data User

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, data); err != nil {
		panic(err)
	}
}

func readdata(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		data.Name = name
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	data.Name = "John Smith"

	t, err := template.ParseFiles("show.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	homeTemplate, err = template.ParseFiles("show.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", readdata)
	http.HandleFunc("/home", home)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":4080", nil); err != nil {
		log.Fatal(err)
	}
}
