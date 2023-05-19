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
