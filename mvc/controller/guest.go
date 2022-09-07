package controller

import (
	"fmt"
	"net/http"
)

type Guest struct {
}

func (m Guest) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "guest index")
}
