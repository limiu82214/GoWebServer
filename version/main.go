package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):]
	r.ParseForm()
	str := r.Form.Get("str")

	t, err := template.ParseFiles("version/" + file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = t.Execute(w, str)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
}
