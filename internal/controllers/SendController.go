// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// SendDataHandle @title SendDataHandle
// @description 发送单条数据到飞书服务器。
// @description /send/密码
// @param w http.ResponseWriter
// @param req *http.Request.
func SendDataHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "请使用POST方法！")
		return
	}

	parameter := strings.Split(strings.TrimPrefix(req.URL.Path, "/send/"), "/")
	if len(parameter) < 1 {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "请输入密码！")
		return
	}

	if !feiShuData.CheckPassWord(parameter[0]) {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "密码错误！")
		return
	}
	feiShuConfig.GetTenantAccessToken()

	// 请求飞书服务器
	request, _ := http.NewRequest("POST", "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/"+feiShuConfig.SpreadsheetToken+"/values_append", req.Body)
	defer request.Body.Close()
	request.Header.Add("Authorization", "Bearer "+feiShuConfig.Token)
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(request)

	// 返回请求飞书服务器的结果
	defer response.Body.Close()
	w.WriteHeader(response.StatusCode)
	readAll, _ := io.ReadAll(response.Body)
	w.Write(readAll)
}

// SendDatasHandle @title SendDatasHandle
// @description 发送多条数据到飞书服务器。
// @description /sends/工作表Id/密码
// @param w http.ResponseWriter
// @param req *http.Request.
func SendDatasHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "请使用POST方法！")
		return
	}

	parameter := strings.Split(strings.TrimPrefix(req.URL.Path, "/sends/"), "/")
	if len(parameter) < 2 {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "参数错误！")
		return
	}

	if !feiShuData.CheckPassWord(parameter[1]) {
		w.WriteHeader(http.StatusPaymentRequired)
		io.WriteString(w, "密码错误！")
		return
	}

	reqData, _ := io.ReadAll(req.Body)
	var values []interface{}
	json.Unmarshal(reqData, &values)

	pushData.ValueRange.Range = parameter[0]
	pushData.ValueRange.Values = append(pushData.ValueRange.Values, values)

	sendCount := 10
	if len(pushData.ValueRange.Values) < sendCount {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "已记录"+strconv.Itoa(len(pushData.ValueRange.Values))+"条数据")
		return
	}

	feiShuConfig.GetTenantAccessToken()

	marshal, _ := json.Marshal(pushData)
	fmt.Println(string(marshal))
	// 请求飞书服务器
	request, _ := http.NewRequest("POST", "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/"+feiShuConfig.SpreadsheetToken+"/values_append", bytes.NewReader(marshal))
	defer request.Body.Close()
	request.Header.Add("Authorization", "Bearer "+feiShuConfig.Token)
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(request)
	if response.StatusCode == http.StatusOK {
		pushData.ValueRange.Values = pushData.ValueRange.Values[sendCount:]
	}

	// 返回请求飞书服务器的结果
	defer response.Body.Close()
	w.WriteHeader(response.StatusCode)
	readAll, _ := io.ReadAll(response.Body)
	w.Write(readAll)
}
