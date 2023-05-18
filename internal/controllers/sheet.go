package controllers

import (
	"io"
	"net/http"
	"strings"
)

func QuerySheetsInfo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	parameter := strings.Split(strings.TrimPrefix(req.URL.Path, "/sheets/"), "/")
	if len(parameter) < 1 {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "参数错误！")
		return
	}

	if !feiShuData.CheckPassWord(parameter[0]) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "密码错误。")
		return
	}

	//...
}
