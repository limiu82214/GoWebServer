package controller

import (
	"net/http"
)

type Guest struct {
}

func (g Guest) Index(w http.ResponseWriter, r *http.Request) (rst string, err error) {
	// 使用 model
	// 取會員資料
	userID := "guest"

	// ------------------
	rst = userID
	err = nil
	return
}
