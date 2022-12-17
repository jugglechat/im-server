package main

import (
	"encoding/binary"
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"google.golang.org/protobuf/proto"
)

/*
int first = buf.readByte();
        int second = buf.readByte();
        int digit;
        int code = first;
        int msgLength = 0;
        int multiplier = 1;
        int lengthSize = 0;
        do {
            lengthSize++;
            digit = buf.readByte();
            code = code ^ digit;
            msgLength += (digit & 0x7f) * multiplier;
            multiplier *= 128;
            if ((digit & 0x80) > 0 && !buf.isReadable()) {
                resumeTimer(ctx);
                buf.resetReaderIndex();
                return null;
            }
        } while ((digit & 0x80) > 0);
        if (code != second) {
            close(ctx, buf);
            return null;
        }
*/

type MyActor struct {
}

func (a *MyActor) OnReceive(msg proto.Message) {
	fmt.Println("msg")
}

func main() {
	// TestNetty()
}

func main_a() {
	actor := clusters.BaseProcessActor(&MyActor{})

	receive := actor.(actorsystem.IReceiveHandler)

	msg := &pbobjs.RpcMessageWraper{}

	receive.OnReceive(msg)
}

func TestNetty() {
	// setup child pipeline initializer.
	childInitializer := func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.LengthFieldCodec(binary.LittleEndian, 1024, 0, 2, 0, 2)).
			AddLast(format.TextCodec())
		// AddLast(codec.ImCodecHandler{}).
		// AddLast(EchoHandler{
		// 	role: "Server",
		// })
	}

	// setup client pipeline initializer.
	clientInitializer := func(channel netty.Channel) {
		channel.Pipeline().
			// AddLast(frame.LengthFieldCodec(binary.LittleEndian, 1024, 0, 2, 0, 2)).
			// AddLast(format.TextCodec()).
			AddLast(codec.ImCodecHandler{}).
			AddLast(EchoHandler{
				role: "Client",
				flag: true,
			})
	}

	// new bootstrap
	var bootstrap = netty.NewBootstrap(netty.WithChildInitializer(childInitializer), netty.WithClientInitializer(clientInitializer))

	// connect to the server after 1 second
	// time.AfterFunc(time.Second, func() {
	// 	_, err := bootstrap.Connect("127.0.0.1:6565", nil)
	// 	utils.Assert(err)

	// })

	// setup bootstrap & startup server.
	bootstrap.Listen("0.0.0.0:6565").Sync()
}

type EchoHandler struct {
	role string
	flag bool
}

func (l EchoHandler) HandleActive(ctx netty.ActiveContext) {
	if l.flag {
		fmt.Println(l.role, "->", "active:", ctx.Channel().RemoteAddr())
		msgHeader := &codec.MsgHeader{Version: codec.Version_1}
		msg := codec.NewConnectMessageWithHeader(msgHeader)
		msg.MsgBody = &codec.ConnectMsgBody{
			ProtoId:  "protoId",
			DeviceId: "deviceId",
		}
		ctx.Write(msg)
	}

	//ctx.Write("Hello I'm " + l.role)

	//ctx.HandleActive()
}

func (l EchoHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if !l.flag {
		fmt.Println("atachemetn:", ctx.Attachment())
		fmt.Println(l.role, "->", "handle read:", message)
		m, ok := message.(*codec.ConnectMessage)
		if ok {
			fmt.Println("name:", m.MsgBody.ProtoId, "\tage:", m.MsgBody.DeviceId)
		} else {
			fmt.Println("xxxxxxx")
		}
	}

	//ctx.HandleRead(message)
}

func (l EchoHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	if l.flag {
		fmt.Println(l.role, "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	}
	//ctx.HandleInactive(ex)
}
