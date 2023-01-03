package main

import (
	"fmt"
	"time"

	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"
	"github.com/jugglechat/im-server/services/connectmanager/clients"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
)

func main() {
	cli := clients.NewImClient("127.0.0.1:9001", "appkey", "CgZhcHBrZXkaIDAr072n8uOcw5YBeKCcQ+QCw4m6YWhgt99U787/dEJS")

	cli.Connect("network", "ispNum", func(code clients.ClientErrorCode, connAck *codec.ConnectAckMsgBody) {
		if code == clients.ClientErrorCode_Success {
			fmt.Println("connect success.")
			SendPrivateMsg(cli)
			//SendMsgTest(cli)
			//QueryTest(cli)
			//PingTest(cli)
		} else {
			fmt.Println("connect fail.", code)
		}
	})
	cli.Disconnect()
	time.Sleep(5 * time.Second)
}
func SendPrivateMsg(cli *clients.ImClient) {
	if cli != nil {
		upMsg := pbobjs.UpMsg{
			MsgType:    "txtMsg",
			MsgContent: []byte(`{"content":"msg content"}`),
			Flags:      1,
		}
		data, _ := tools.PbMarshal(&upMsg)
		cli.Publish("pMsg", "tarId", data, func(code clients.ClientErrorCode, pubAck *codec.PublishAckMsgBody) {
			fmt.Println(code)
		})
	}
}
func PingTest(cli *clients.ImClient) {
	if cli != nil {
		for i := 0; i < 3; i++ {
			cli.Ping(func(code clients.ClientErrorCode) {
				fmt.Println("pong", code)
			})
			time.Sleep(1 * time.Second)
		}
	}
}
func QueryTest(cli *clients.ImClient) {
	if cli != nil {
		for i := 0; i < 10; i++ {
			cli.Query("queryMethod", "queryTarget", []byte{1, 2, 3, 4, 5}, func(code clients.ClientErrorCode, qryAck *codec.QueryAckMsgBody) {
				if code == clients.ClientErrorCode_Success {
					fmt.Println("code:", code, "\tdata:", qryAck.Data)
				}
			})
			time.Sleep(1 * time.Second)
		}
	}
}
func SendMsgTest(cli *clients.ImClient) {
	if cli != nil {
		for i := 0; i < 10; i++ {
			cli.Publish("sendMsg", "targetId", []byte{1, 2, 3}, func(code clients.ClientErrorCode, pubAck *codec.PublishAckMsgBody) {
				if code == clients.ClientErrorCode_Success {
					fmt.Println("code:", code, "\tmsgId", pubAck.MsgId, "\tmsgTime:", pubAck.Timestamp)
				}
			})
			time.Sleep(1 * time.Second)
		}
	}
}
