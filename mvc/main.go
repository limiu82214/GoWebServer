package main

import (
	"errors"
	"fmt"
	"log"
	"mvc/controller"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", handle)
	// http.HandleFunc("/", controller.NotFound404)
	// http.HandleFunc("/member/", handle)
	// http.HandleFunc("/guest/", ghandle)
	log.Fatal(http.ListenAndServe(Host, nil))

}
func getControllerAction(path string) (ctl string, action string, err error) {
	ctl = ""
	action = ""
	err = nil
	p := strings.Split(path, "/")
	if len(p) <= 2 {
		err = errors.New("path < 2")
		return
	}
	// 拆出 controller 與 action
	ctl = p[1]
	action = p[2]
	action = strings.ToUpper(action)[0:1] + action[1:]
	return
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctl, action, err := getControllerAction(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}

	controller.Router.AddController(controller.Index)
	controller.Router.AddController(controller.Member)
	controller.Router.AddController(controller.Guest)

	if controller.Router.IsNotActionExist(ctl, action) {
		ctl, action = "index", "NotFound404"
	}

	data, err := controller.Router.Execute(ctl, action, r)
	if err != nil {
		log.Fatal(err)
	}

	viewPath := RelativeDIR + "view/" + ctl + "/" + action + ".html"
	template, err := template.ParseFiles(viewPath)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	_ = template.Execute(w, data)
}
