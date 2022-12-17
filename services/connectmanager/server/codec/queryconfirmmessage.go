package codec

import "github.com/jugglechat/im-server/commons/tools"

type QueryConfirmMessage struct {
	MsgHeader
	MsgBody *QueryConfirmMsgBody
}

func NewQueryConfirmMessageWithHeader(header *MsgHeader) *QueryConfirmMessage {
	msg := &QueryConfirmMessage{
		MsgHeader: MsgHeader{
			Version:     Version_1,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	return msg
}

func (msg *QueryConfirmMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *QueryConfirmMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &QueryConfirmMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
