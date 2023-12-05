package common

import "net/http"

type Result struct {
	Code int         `json:"code" `
	Msg  string      `json:"msg" `
	Data interface{} `json:"data" `
}

type PageResult struct {
	Code     int         `json:"code" `
	Msg      string      `json:"msg" `
	Data     interface{} `json:"data" `
	PageInfo PageInfo    `json:"pageInfo" `
}

func Success(data interface{}) (result *Result) {
	return &Result{
		Code: http.StatusOK,
		Msg:  "请求成功",
		Data: data,
	}
}
func PageSuccess(data interface{}, pageInfo PageInfo) (result *PageResult) {
	return &PageResult{
		Code:     http.StatusOK,
		Msg:      "请求成功",
		Data:     data,
		PageInfo: pageInfo,
	}
}

func Fail(data interface{}, err error) (result *Result) {
	return &Result{
		Code: ErrorCode,
		Msg:  err.Error(),
		Data: data,
	}
}
func DbFail(data interface{}) (result *Result) {
	return &Result{
		Code: ErrorCode,
		Data: data,
	}
}
