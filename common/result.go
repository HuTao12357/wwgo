package common

import "net/http"

type Result struct {
	Code int         `json:"code" `
	Msg  string      `json:"msg" `
	Data interface{} `json:"data" `
}

func Success(data interface{}) (result *Result) {
	return &Result{
		Code: http.StatusOK,
		Msg:  "请求成功",
		Data: data,
	}
}