package dbs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AppTable struct {
	ID           int64 `gorm:"primary_key"`
	AppKey       string
	AppSecret    string
	AppSecureKey string
	AppStatus    string
	CreatedTime  time.Time
}

func (app AppTable) TableName() string {
	return "im_apps"
}

func (app AppTable) FindByAppkey(appkey string) *AppTable {
	var appItem AppTable
	err := db.Where("app_key=?", appkey).Take(&appItem).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &appItem
}

type AppExtTable struct {
	ID           int64 `gorm:"primary_key"`
	AppKey       string
	AppItemKey   string
	AppItemValue string
	CreatedTime  time.Time
}

func (appExt AppExtTable) TableName() string {
	return "im_appexts"
}

func (appExt AppExtTable) FindListByAppkey(appkey string) []*AppExtTable {
	var list []*AppExtTable
	db.Where("app_key=?", appkey).Find(&list)
	return list
}
