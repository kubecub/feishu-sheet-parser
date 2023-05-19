package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SetFeiShuConfig @title SetFeiShuConfig
// @description 设置飞书配置
// @description /set/表格Id/应用Id/应用Secret
// @param w http.ResponseWriter
// @param req *http.Request
func SetFeiShuConfig(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	parameter := strings.Split(strings.TrimPrefix(req.URL.Path, "/set/"), "/")
	if len(parameter) < 3 {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "参数错误！")
		return
	}

	if feiShuConfig.CheckIsCreate() {
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "已经设置过了。\n如果需要修改请通过更新的方式。")
		return
	}

	feiShuConfig.SpreadsheetToken = parameter[0]
	feiShuConfig.AppId = parameter[1]
	feiShuConfig.AppSecret = parameter[2]
	feiShuConfig.GetTenantAccessToken()

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("配置成功！\n密码：%s\n", feiShuData.BuildPassWord()))
}

// UpdateFeiShuConfig @title UpdateFeiShuConfig
// @description 更新飞书配置
// @description /update/表格Id/应用Id/应用Secret
// @param w http.ResponseWriter
// @param req *http.Request
func UpdateFeiShuConfig(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	parameter := strings.Split(strings.TrimPrefix(req.URL.Path, "/update/"), "/")
	if len(parameter) < 4 {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "参数错误！")
		return
	}

	if !feiShuData.CheckPassWord(parameter[3]) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "密码错误。")
		return
	}

	feiShuConfig.SpreadsheetToken = parameter[0]
	feiShuConfig.AppId = parameter[1]
	feiShuConfig.AppSecret = parameter[2]

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "更新成功。")
}
