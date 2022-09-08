package controller

import (
	"net/http"
)

var Guest *guest

func init() {
	Guest = &guest{}
	Guest.name = "guest"
}

type guest struct{ ControllerBase }

func (g *guest) Index(r *http.Request) (rst string, err error) {
	// 使用 model
	// 取會員資料
	userID := "guest"

	// ------------------
	rst = userID
	err = nil
	return
}
