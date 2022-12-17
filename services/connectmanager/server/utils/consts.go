package utils

const (
	Platform_iOS     string = "iOS"
	Platform_Android string = "Android"
	Platform_Web     string = "Web"
	Platform_PC      string = "PC"
)

type ActionType int

const (
	Action_Query        ActionType = 0
	Action_QueryAck     ActionType = 1
	Action_QueryConfirm ActionType = 2

	Action_UserPub    ActionType = 3
	Action_UserPubAck ActionType = 4

	Action_ServerPub    ActionType = 5
	Action_ServerPubAck ActionType = 6

	Action_Connect    ActionType = 7
	Action_Disconnect ActionType = 8
	Action_ConnectErr ActionType = 9
)

const (
	StateKey_ObfuscationCode      string = "state.obfuscation_code"
	StateKey_CtxLocker            string = "state.ctx_locker"
	StateKey_ServerPubCallbackMap string = "state.pub_callback_map"
	StateKey_QueryConfirmMap      string = "state.qry_confirm_map"
	StateKey_ConnectSession       string = "state.connect_session"
	StateKey_ConnectCreateTime    string = "state.connect_timestamp"
	StateKey_ServerMsgIndex       string = "state.server_msg_index"
	StateKey_ClientMsgIndex       string = "state.client_msg_index"
	StateKey_Appkey               string = "state.appkey"
	StateKey_UserID               string = "state.userid"
	StateKey_Platform             string = "state.platform"
	StateKey_Version              string = "state.version"
	StateKey_ClientIp             string = "state.client_ip"
)
const (
	ConnectAckState_Access              int32 = 0
	ConnectAckState_Redirect            int32 = 1
	ConnectAckState_UnsupportedPlatform int32 = 2
	ConnectAckState_AuthorizeFailed     int32 = 3
	ConnectAckState_AuthorizeExpired    int32 = 4
	ConnectAckState_AppBlock            int32 = 5
	ConnectAckState_UserBlock           int32 = 6
)

const (
	DisconnectState_None       int32 = 0
	DisconnectState_Redirect   int32 = 1
	DisconnectState_OtherLogin int32 = 2
	DisconnectState_Quit       int32 = 3
	DisconnectState_Logout     int32 = 4
	DisconnectState_UserBlock  int32 = 6
)

type CloseReason int

const (
	Close_Normal CloseReason = 0
)
