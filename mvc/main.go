package main

import (
	"log"
	"mvc/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.NotFound404)

	m := controller.Member{}
	http.HandleFunc("/member/", m.Index)

	g := controller.Guest{}
	http.HandleFunc("/guest/", g.Index)
	log.Fatal(http.ListenAndServe(Host, nil))
}
