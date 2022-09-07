package controller

import (
	"fmt"
	"net/http"
)

type Member struct {
}

func (m Member) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "member index")
}
