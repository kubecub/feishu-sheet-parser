package model

import (
	"github.com/cubxxw/feishu-sheet-parser/internal/utils"
)

type FeiShuData struct {
	password string
}

func (feiShuData *FeiShuData) BuildPassWord() string {
	feiShuData.password = utils.RandStr(128)
	return feiShuData.password
}

func (feiShuData *FeiShuData) CheckPassWord(password string) bool {
	return feiShuData.password != "" && feiShuData.password == password
}