package service

import (
	"errors"
	"goChat/db"
	"goChat/model"
	"goChat/util"
	"time"
)

func AddFriend(oid, tid int64) error {
	if oid == tid {
		return errors.New(util.IText.CAN_NOT_TAOWA)
	}
	tem := model.Contact{}
	_, _ = db.Engine.Where("owner = ?", oid).And("target = ?", tid).And("type = ?", model.CONTACT_TYPE_USER).Get(&tem)
	//存在好友关系
	if tem.Id > 0 {
		return errors.New(util.IText.HAS_BEING_FRIEND)
	}
	//使用事务
	session := db.Engine.NewSession()
	_ = session.Begin()
	_, e1 := session.InsertOne(model.Contact{
		Owner:    oid,
		Target:   tid,
		CreateAt: time.Now(),
		Type:     model.CONTACT_TYPE_USER,
	})
	_, e2 := session.InsertOne(model.Contact{
		Owner:    tid,
		Target:   oid,
		CreateAt: time.Now(),
		Type:     model.CONTACT_TYPE_USER,
	})
	if e1 == nil && e2 == nil {
		e3 := session.Commit()
		return e3
	} else {
		if e1 != nil {
			return e1
		} else {
			return e2
		}
	}
}
func RemoveFriendById(uid int64) {

}
func AddCommunityById(cid int64) {

}
func RemoveCommunityById(cid int64) {

}
func SearchFriendById(uid int64) {

}
func SearchCommunityById(cid int64) {

}
func CreateCommunity() {

}
