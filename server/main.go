package main

import (
	"encoding/json"
	"fmt"
	"goChat/db"
	"net/http"
)

//字段首字母一定要大写

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func main() {
	//请求接收
	http.HandleFunc("/user/login", UserLogin)
	//启动web服务器
	_ = http.ListenAndServe(":8090", nil)
	db.InitDb()

}

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = request.ParseForm()
	phoneNumber := request.PostForm.Get("phoneNumber")
	password := request.PostForm.Get("password")
	if phoneNumber == "18328023227" && password == "123456" {
		obj := make(map[string]interface{})
		obj["id"] = 1
		obj["token"] = "yes"
		Resp(writer, 500, obj, "登录成功")
	} else {
		Resp(writer, 500, nil, "用户名或密码错误")
	}
}
func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	result := BaseResponse{
		Data: data,
		Code: code,
		Msg:  msg,
	}
	resultStrJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = writer.Write(resultStrJson)
}

func RegisterView() {

}
