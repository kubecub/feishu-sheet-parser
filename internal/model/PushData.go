// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package model

type PushData struct {
	ValueRange struct {
		Range  string          `json:"range"`
		Values [][]interface{} `json:"values"`
	} `json:"valueRange"`
}

type Value struct {
	Value interface{}
}
