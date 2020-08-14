package model

import "time"

type Contact struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Owner    int64     `xorm:"bigint(20)" form:"owner" json:"owner"`
	Target   int64     `xorm:"bigint(20)" form:"target" json:"target"`
	Type     int64     `xorm:"bigint(20)" form:"type" json:"type"`
	Motto    string    `xorm:"varchar(256)" form:"motto" json:"motto"`
	CreateAt time.Time `xorm:"datetime" form:"createat" json:"createat"`
}

const (
	CONTACT_TYPE_USER      = 1
	CONTACT_TYPE_COMMUNITY = 2
)
