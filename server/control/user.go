package control

import (
	"goChat/service"
	"goChat/util"
	"net/http"
)

var text = util.I18n("zh")

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	user, err := service.Login(mobile, password)
	if err != nil {
		util.Resp(writer, 500, nil, err.Error())
	} else {
		util.Resp(writer, 200, user, text.LOGIN_SUCCESS)
	}
}
func UserRegister(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	user, err := service.Register(mobile, password)
	if err != nil {
		util.Resp(writer, 500, nil, err.Error())
	} else {
		util.Resp(writer, 200, user, text.REGISTER_SUCCESS)
	}
}
