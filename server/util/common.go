package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//字段 首字母一定要大写
type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
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

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePassword(plainPassword, salt, password string) bool {
	return Md5Encode(plainPassword+salt) == password
}
func MakePassword(plainPassword, salt string) string {
	return Md5Encode(plainPassword + salt)
}
