package main

import (
	"fmt"
	"goChat/control"
	"goChat/db"
	"goChat/util"
	"io/ioutil"
	"net/http"
)

const (
	PORT = ":9090"
)

//请求拦截
func handleIntercept(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		w.WriteHeader(http.StatusOK)
		userToken := r.Header.Get("token")
		//是否有token
		if userToken == "" {
			util.Resp(w, 500, nil, util.IText.PERMISSION_DENY)
		} else {
			_, err := util.ParseToken(userToken)
			//是否过期
			if err == nil {
				h(w, r)
			} else {
				util.Resp(w, 500, nil, util.IText.TOKEN_EXPIRED)
			}
		}
	}
}

//接口初始化统一管理
func ApiInit() {
	http.HandleFunc("/user/login", control.UserLogin)
	http.HandleFunc("/user/register", control.UserRegister)
	http.HandleFunc("/user/test", testDelete)
	http.HandleFunc("/user/token/check", handleIntercept(control.CheckUserToken))
}

func testDelete(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE")

	req, _ := ioutil.ReadAll(request.Body)
	fmt.Printf(string(req))
	_, _ = w.Write(req)
}

//主函数
func main() {
	fmt.Println(util.MakePassword("olive888", "0r4p09kxq6cj"))
	//连接数据库
	db.InitDb()

	//接口初始化
	ApiInit()

	//启动web服务器
	_ = http.ListenAndServe(PORT, nil)

}
