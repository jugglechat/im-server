package codec

import "github.com/jugglechat/im-server/commons/tools"

type DisconnectMessage struct {
	MsgHeader
	MsgBody *DisconnectMsgBody
}

func NewDisconnectMessage(msgBody *DisconnectMsgBody) *DisconnectMessage {
	msg := &DisconnectMessage{
		MsgHeader: MsgHeader{
			Version: Version_1,
		},
		MsgBody: msgBody,
	}
	msg.SetCmd(Cmd_Disconnect)
	msg.SetQoS(QoS_NoAck)
	return msg
}
func NewDisconnectMessageWithHeader(header *MsgHeader) *DisconnectMessage {
	msg := &DisconnectMessage{
		MsgHeader: MsgHeader{
			Version:     Version_1,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	msg.SetCmd(Cmd_Disconnect)
	msg.SetQoS(QoS_NoAck)
	return msg
}

func (msg *DisconnectMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *DisconnectMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &DisconnectMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
