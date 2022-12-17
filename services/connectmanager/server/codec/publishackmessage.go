package codec

import "github.com/jugglechat/im-server/commons/tools"

type ServerPublishAckMessage struct {
	MsgHeader
	MsgBody *PublishAckMsgBody
}

type UserPublishAckMessage struct {
	MsgHeader
	MsgBody *PublishAckMsgBody
}

func NewServerPublishAckMessage(msgBody *PublishAckMsgBody) *ServerPublishAckMessage {
	msg := &ServerPublishAckMessage{
		MsgHeader: MsgHeader{
			Version: Version_1,
		},
		MsgBody: msgBody,
	}
	msg.SetCmd(Cmd_PublishAck)
	msg.SetQoS(QoS_NoAck)
	return msg
}
func NewUserPublishAckMessage(msgBody *PublishAckMsgBody) *UserPublishAckMessage {
	msg := &UserPublishAckMessage{
		MsgHeader: MsgHeader{
			Version: Version_1,
		},
		MsgBody: msgBody,
	}
	msg.SetCmd(Cmd_PublishAck)
	msg.SetQoS(QoS_NoAck)
	return msg
}

func NewUserPublishAckMessageWithHeader(header *MsgHeader) *UserPublishAckMessage {
	msg := &UserPublishAckMessage{
		MsgHeader: MsgHeader{
			Version:     Version_1,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	return msg
}

func (msg *UserPublishAckMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *UserPublishAckMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &PublishAckMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}

func NewServerPublishAckMessageWithHeader(header *MsgHeader) *ServerPublishAckMessage {
	msg := &ServerPublishAckMessage{
		MsgHeader: MsgHeader{
			Version:     Version_1,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	return msg
}

func (msg *ServerPublishAckMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *ServerPublishAckMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &PublishAckMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
