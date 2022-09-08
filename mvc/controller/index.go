package controller

import (
	"net/http"
)

var Index *index

func init() {
	Index = &index{}
	Index.name = "index"
}

type index struct{ ControllerBase }

func (i *index) NotFound404(r *http.Request) (rst string, err error) {
	rst = "404"
	err = nil
	return
}
