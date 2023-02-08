package utils

import (
	"encoding/json"
	"net/http"
)

// response 返回定义的结构体
type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		panic("json格式转换错误")
	}
	w.Write(ret)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Data:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		panic("json格式转换错误")
	}
	w.Write(ret)
}

// 失败的响应
func RespFial(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

// 成功的响应
func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func RespOkList(w http.ResponseWriter, data interface{}, total int) {
	RespList(w, 0, data, total)
}
