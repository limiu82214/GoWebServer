package controller

import (
	"fmt"
	"net/http"
)

func NotFound404(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404")
}
