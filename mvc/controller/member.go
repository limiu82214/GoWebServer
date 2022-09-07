package controller

import (
	"net/http"
)

type Member struct {
}

func (m Member) Index(w http.ResponseWriter, r *http.Request) (rst string, err error) {
	// 使用 model
	// 取會員資料
	userID := "lim"

	// ------------------
	rst = userID
	err = nil
	return
}
