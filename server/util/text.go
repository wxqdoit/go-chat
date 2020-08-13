package util

type Text struct {
	USER_EXIST       string
	USER_NOT_EXIST   string
	USER_ERROR       string
	LOGIN_SUCCESS    string
	REGISTER_SUCCESS string
	PERMISSION_DENY  string
}

var zh *Text = &Text{
	USER_EXIST:       "该用户已存在",
	USER_NOT_EXIST:   "用户不存在",
	USER_ERROR:       "用户名或密码错误",
	LOGIN_SUCCESS:    "登录成功",
	REGISTER_SUCCESS: "注册成功",
	PERMISSION_DENY:  "无权限",
}
var en *Text = &Text{
	USER_EXIST:       "User existed!",
	USER_NOT_EXIST:   "User not exist",
	USER_ERROR:       "Username or password error",
	LOGIN_SUCCESS:    "login success",
	REGISTER_SUCCESS: "register success",
	PERMISSION_DENY:  "permission deny",
}

func I18n(lang string) (z *Text) {
	switch lang {
	case "zh":
		return zh
	case "en":
		return en
	default:
		return zh
	}
}

var IText *Text = I18n("zh")
