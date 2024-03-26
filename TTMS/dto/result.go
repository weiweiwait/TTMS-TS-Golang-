package dto

import "TTMS/pkg/e"

type Result struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func Fail(code int, err error) *Result {
	result := &Result{
		Code: code,
		Msg:  e.GetMsg(code),
	}
	if err != nil {
		result.Data = err.Error()
	}
	return result
}

func Success(code int, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}
