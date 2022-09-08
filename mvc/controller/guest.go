package controller

import (
	"net/http"
)

func init() {

}

type Guest struct {
	ControllerBase
}

func NewGuest() *Guest {
	g := &Guest{}
	g.name = "guest"
	return g
}

func (g *Guest) Index(w http.ResponseWriter, r *http.Request) (rst string, err error) {
	// 使用 model
	// 取會員資料
	userID := "guest"

	// ------------------
	rst = userID
	err = nil
	return
}
