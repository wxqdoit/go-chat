package model

import "time"

type Community struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Name     string    `xorm:"varchar(256)" form:"name" json:"name"`
	Motto    string    `xorm:"varchar(256)" form:"motto" json:"motto"`
	Owner    int64     `xorm:"bigint(64)" form:"owner" json:"owner"`
	Icon     string    `xorm:"varchar(256)" form:"icon" json:"icon"`
	Type     int64     `xorm:"bigint(20)" form:"type" json:"type"`
	CreateAT time.Time `xorm:"datetime" form:"createat" json:"createat"`
}

const (
	COMMUNITY_TYPE_COM = 1
)
