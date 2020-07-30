package _http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data string `json:"data"`
}

func Out(w http.ResponseWriter, msg string, code int, data [][]string) {
	bytes, _ := json.Marshal(data)
	jsonStr := string(bytes)
	resp := Response{msg, code, jsonStr}
	_resp, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println(err)
	}
	var resp_str string = string(_resp)
	fmt.Fprintf(w, resp_str)
}
