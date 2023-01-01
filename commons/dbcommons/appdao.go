package dbcommons

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AppTable struct {
	ID           int64     `gorm:"primary_key"`
	AppKey       string    `gorm:"app_key"`
	AppSecret    string    `gorm:"app_secret"`
	AppSecureKey string    `gorm:"app_secure_key"`
	AppStatus    string    `gorm:"app_status"`
	CreatedTime  time.Time `gorm:"created_time"`
}

func (app AppTable) TableName() string {
	return "apps"
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
	return "appexts"
}

func (appExt AppExtTable) FindListByAppkey(appkey string) []*AppExtTable {
	var list []*AppExtTable
	db.Where("app_key=?", appkey).Find(&list)
	return list
}
