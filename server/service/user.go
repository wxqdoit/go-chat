package service

import (
	"errors"
	"fmt"
	"goChat/db"
	"goChat/model"
	"goChat/util"
	"math/rand"
	"time"
)

var text = util.I18n("zh")

func Register(mobile string, password string) (user model.User, err error) {
	_user := model.User{}
	_, err = db.Engine.Where("mobile=?", mobile).Get(&_user)
	if err != nil {
		return _user, err
	}
	if _user.Id > 0 {
		return _user, errors.New(text.USER_EXIST)
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	_user.Password = util.MakePassword(password, salt)
	_user.Mobile = mobile
	_user.Salt = salt
	_user.CreateAt = time.Now()

	_, err = db.Engine.InsertOne(&_user)
	if err == nil {
		return _user, nil
	} else {
		return _user, err
	}
}

func Login(mobile string, password string) (user model.User, err error) {
	_user := model.User{}
	_, err = db.Engine.Where("mobile=? ", mobile).Get(&_user)
	if err != nil {
		return _user, err
	}
	if _user.Id > 0 {
		if util.ValidatePassword(password, _user.Salt, _user.Password) {
			return _user, nil
		} else {
			return _user, errors.New(text.USER_ERROR)
		}
	} else {
		return _user, errors.New(text.USER_NOT_EXIST)
	}
}
