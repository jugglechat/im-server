package dbs

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jugglechat/im-server/commons/dbcommons"
)

type UserDao struct {
	UserId       string    `gorm:"user_id"`
	Nickname     string    `gorm:"nickname"`
	UserPortrait string    `gorm:"user_portrait"`
	CreatedTime  time.Time `gorm:"created_time"`
}

func (user UserDao) TableName() string {
	return "users"
}
func (user UserDao) FindByUserId(userId string) *UserDao {
	var item UserDao
	err := dbcommons.GetDb().Where("user_id=?", userId).Take(&item).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &item
}
func (user UserDao) Create(item UserDao) error {
	err := dbcommons.GetDb().Create(&item).Error
	return err
}
func (user UserDao) CreateOrUpdate(item UserDao) error {
	return dbcommons.GetDb().Exec(fmt.Sprintf("INSERT INTO %s (user_id,nickname,user_portrait,created_time)VALUES(?,?,?,?) ON DUPLICATE KEY UPDATE nickname=?,user_portrait=?,created_time=?", user.TableName()), item.UserId, item.Nickname, item.UserPortrait, time.Now(), item.Nickname, item.UserPortrait, time.Now()).Error
}

func (user UserDao) UpdateLogo(id int64, url string) error {
	upd := map[string]interface{}{}
	upd["logo_url"] = url
	return dbcommons.GetDb().Model(&user).Where("id=?", id).Update(upd).Error
}
