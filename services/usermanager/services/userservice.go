package services

import (
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/services/usermanager/dbs"
)

func AddUser(ctx clusters.BaseContext, userId, nickname, userPortrait string) {
	dao := dbs.UserDao{}
	dao.CreateOrUpdate(dbs.UserDao{
		UserId:       userId,
		Nickname:     nickname,
		UserPortrait: userPortrait,
	})
}
