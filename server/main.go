package main

import (
	"fmt"
	"goChat/control"
	"goChat/db"
	"goChat/util"
	"net/http"
)

//请求拦截
func handleIntercept(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		userToken := r.Header.Get("token")
		if userToken == "" {
			util.Resp(w, 500, nil, util.IText.PERMISSION_DENY)
		} else {
			c, err := util.ParseToken(userToken)
			fmt.Println("----userToken-----")
			fmt.Print(c, err)
			fmt.Println("----userToken-----")
			if err == nil {
				h(w, r)
			} else {
				util.Resp(w, 500, nil, error.Error(err))
			}
		}
	}
}

//接口初始化统一管理
func ApiInit() {
	http.HandleFunc("/user/login", control.UserLogin)
	http.HandleFunc("/user/register", control.UserRegister)
	http.HandleFunc("/user/token/check", handleIntercept(control.CheckUserToken))
}

//主函数
func main() {

	//连接数据库
	db.InitDb()

	//接口初始化
	ApiInit()

	//启动web服务器
	_ = http.ListenAndServe(":8090", nil)

}
