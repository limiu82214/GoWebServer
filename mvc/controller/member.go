package controller

import (
	"net/http"
)

var Member *member

func init() {
	Member = &member{}
	Member.name = "member"
}

type member struct{ ControllerBase }

func (m *member) Index(r *http.Request) (rst string, err error) {
	// 使用 model
	// 取會員資料
	userID := "lim2"

	// ------------------
	rst = userID
	err = nil
	return
}
