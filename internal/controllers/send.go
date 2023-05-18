package controllers

import (
	"feishu_sheets_api/model"
	"net/http"
)

var (
	feiShuConfig = &model.FeiShuConfig{}
	feiShuData   = &model.FeiShuData{}
	pushData     = &model.PushData{}
)

func SendDataHandle(w http.ResponseWriter, req *http.Request) {

}

func SendDatasHandle(w http.ResponseWriter, req *http.Request) {

}
