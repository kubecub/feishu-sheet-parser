package controllers

import (
	"io"
	"net/http"
	"strings"
)

// QuerySheetsInfo @title QuerySheetsInfo
// @description 获取飞书表格信息。
// @description /sheets/密码
// @param w http.ResponseWriter
// @param req *http.Request
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

	url := "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/" + feiShuConfig.SpreadsheetToken + "/metainfo"

	client := &http.Client{}
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+feiShuConfig.Token)
	w.Header().Set("Content-Type", "application/json")

	response, _ := client.Do(req)
	defer response.Body.Close()
	w.WriteHeader(response.StatusCode)
	readAll, _ := io.ReadAll(response.Body)
	w.Write(readAll)
}
