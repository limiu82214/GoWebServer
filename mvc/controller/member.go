package controller

import (
	"net/http"
)

func init() {

}

type Member struct {
	ControllerBase
}

func NewMember() *Member {
	m := &Member{}
	m.name = "member"
	return m
}

func (m *Member) Index(w http.ResponseWriter, r *http.Request) (rst string, err error) {
	// 使用 model
	// 取會員資料
	userID := "lim2"

	// ------------------
	rst = userID
	err = nil
	return
}
