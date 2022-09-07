package main

import (
	"fmt"
	"log"
	"mvc/controller"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	http.HandleFunc("/", handle)
	// http.HandleFunc("/", controller.NotFound404)
	// http.HandleFunc("/member/", handle)
	// http.HandleFunc("/guest/", ghandle)
	log.Fatal(http.ListenAndServe(Host, nil))

}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", r.URL.Path)
	// 檢查有無 / controller / action /
	// 若無 跳404 return
	// ---------------------------------------
	// 拆出 controller 與 action
	ctl := "member"
	action := "index"
	str := r.URL.Path[len("/"+ctl+"/"):] // 拆action
	fmt.Fprintf(w, "%v\n", str)
	str = strings.ToUpper(str)[0:1] + str[1:]
	fmt.Println(str)

	t := reflect.TypeOf(&controller.Member{}) // ---------------- 怎麼用ctl來TypeOf
	_, isExist := t.MethodByName(str)
	fmt.Println(isExist)
	if !isExist {
		controller.NotFound404(w, r)
	}

	// m := controller.Member{}
	// rst, err := m.Index(w, r)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// path := "view/member/index.html"
	// t, err := template.ParseFiles(RelativeDIR + path)
	// if err != nil {
	// 	fmt.Fprintf(w, "%v", err)
	// 	return
	// }
	// _ = t.Execute(w, rst)
}
