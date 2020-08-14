package util

type Text struct {
	USER_EXIST       string
	USER_NOT_EXIST   string
	USER_ERROR       string
	LOGIN_SUCCESS    string
	REGISTER_SUCCESS string
	PERMISSION_DENY  string
	TOKEN_EXPIRED    string
	CAN_NOT_TAOWA    string
	HAS_BEING_FRIEND string
}

var zh *Text = &Text{
	USER_EXIST:       "该用户已存在",
	USER_NOT_EXIST:   "用户不存在",
	USER_ERROR:       "用户名或密码错误",
	LOGIN_SUCCESS:    "登录成功",
	REGISTER_SUCCESS: "注册成功",
	PERMISSION_DENY:  "无权限",
	TOKEN_EXPIRED:    "登录过期",
	CAN_NOT_TAOWA:    "禁止套娃加自己为好友哦",
	HAS_BEING_FRIEND: "不能当复读机重复添加哟",
}
var en *Text = &Text{
	USER_EXIST:       "User existed!",
	USER_NOT_EXIST:   "User not exist",
	USER_ERROR:       "Username or password error",
	LOGIN_SUCCESS:    "Login success",
	REGISTER_SUCCESS: "Register success",
	PERMISSION_DENY:  "Permission deny",
	TOKEN_EXPIRED:    "Token is expired",
	CAN_NOT_TAOWA:    "Cannot add yourself",
	HAS_BEING_FRIEND: "Already friends",
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
