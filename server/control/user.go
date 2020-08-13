package control

import (
	"github.com/dgrijalva/jwt-go"
	"goChat/service"
	"goChat/util"
	"net/http"
)

var text = util.I18n("zh")

type Claims struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	user, err := service.Login(mobile, password)
	if err != nil {
		util.Resp(writer, 500, nil, err.Error())
	} else {
		str, singErr := util.GenerateToken(mobile, password)
		if singErr == nil {
			user.Token = str
			_, upErr := service.UpdateUserToken(str, mobile)
			if upErr == nil {
				util.Resp(writer, 200, user, text.LOGIN_SUCCESS)
			} else {
				util.Resp(writer, 500, nil, upErr.Error())
			}
		} else {
			util.Resp(writer, 500, nil, singErr.Error())
		}
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
func CheckUserToken(writer http.ResponseWriter, request *http.Request) {
	util.Resp(writer, 200, nil, "")
}
