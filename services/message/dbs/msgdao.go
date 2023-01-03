package dbs

import "github.com/jugglechat/im-server/commons/dbcommons"

type MsgDao struct {
	ID        int64  `gorm:"primary_key"`
	UserId    string `gorm:"user_id"`
	SendTime  int64  `gorm:"send_time"`
	Direction int    `gorm:"direction"`
	MsgBody   []byte `gorm:"msg_body"`
}

func (msg MsgDao) TableName() string {
	return "messages"
}
func (msg MsgDao) Create(item MsgDao) error {
	err := dbcommons.GetDb().Create(&item).Error
	return err
}
func (msg MsgDao) QueryList(userId string, start int64, count int) ([]*MsgDao, error) {
	var items []*MsgDao
	err := dbcommons.GetDb().Where("user_id=? and send_time>?", userId, start).Order("send_time asc").Limit(count).Find(&items).Error
	return items, err
}
