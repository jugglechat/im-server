package services

import (
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"
	"github.com/jugglechat/im-server/services/message/dbs"
)

const (
	MsgDirection_Receive int = 0
	MsgDirection_Send    int = 1
)

func SaveMsg(userId string, msg *pbobjs.DownMsg, direction int) error {
	//save immediately when user online.
	//TODO other wise, use async queue.
	msgBs, _ := tools.PbMarshal(msg)
	msgDao := dbs.MsgDao{}
	err := msgDao.Create(dbs.MsgDao{
		UserId:    userId,
		SendTime:  msg.MsgTime,
		Direction: direction,
		MsgBody:   msgBs,
	})
	return err
}
