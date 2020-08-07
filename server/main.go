package main

import (
	"goChat/control"
	"goChat/db"
	"net/http"
)

//请求拦截
func handleIntercept(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		header := r.Header

		userToken := header.Get("token")
		println(userToken)
		if userToken == "" {
			//util.Resp(w, 500, nil, util.I18n())
		}
		h(w, r)
	}
}

//主函数
func main() {
	// 连接数据库
	db.InitDb()

	//请求接收
	http.HandleFunc("/user/login", handleIntercept(control.UserLogin))
	http.HandleFunc("/user/register", handleIntercept(control.UserRegister))

	//启动web服务器
	_ = http.ListenAndServe(":8090", nil)

}
