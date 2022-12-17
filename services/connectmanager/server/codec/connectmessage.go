package codec

import (
	"github.com/jugglechat/im-server/commons/tools"
)

type ConnectMessage struct {
	MsgHeader
	MsgBody *ConnectMsgBody
}

func NewConnectMessage(msgBody *ConnectMsgBody) *ConnectMessage {
	msg := &ConnectMessage{
		MsgHeader: MsgHeader{
			Version: Version_1,
		},
		MsgBody: msgBody,
	}
	msg.SetCmd(Cmd_Connect)
	msg.SetQoS(QoS_NeedAck)
	return msg
}

func NewConnectMessageWithHeader(header *MsgHeader) *ConnectMessage {
	msg := &ConnectMessage{
		MsgHeader: MsgHeader{
			Version:     Version_1,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	msg.SetCmd(Cmd_Connect)
	msg.SetQoS(QoS_NeedAck)
	return msg
}

func (msg *ConnectMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *ConnectMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &ConnectMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
