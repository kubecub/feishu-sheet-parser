// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package model

import (
	"github.com/kubecub/feishu-sheet-parser/internal/utils"
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
