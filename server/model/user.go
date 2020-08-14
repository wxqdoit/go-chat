package model

import "time"

type User struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Password string    `xorm:"varchar(64)" form:"password" json:"-"`
	Avatar   string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string    `xorm:"varchar(8)" form:"sex" json:"sex"`
	Nickname string    `xorm:"varchar(32)" form:"nickname" json:"nickname"`
	Salt     string    `xorm:"varchar(32)" form:"salt" json:"-"`
	Online   int       `xorm:"int(10)" form:"online" json:"online"`
	Token    string    `xorm:"varchar(256)" form:"token" json:"token"`
	Motto    string    `xorm:"varchar(256)" form:"motto" json:"motto"`
	CreateAt time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
