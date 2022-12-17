package clients

type ClientErrorCode int

const (
	ClientErrorCode_Success        ClientErrorCode = 0
	ClientErrorCode_Unknown        ClientErrorCode = 20000
	ClientErrorCode_SocketFailed   ClientErrorCode = 20001
	ClientErrorCode_ConnectTimeout ClientErrorCode = 20002
	ClientErrorCode_NeedRedirect   ClientErrorCode = 20003
	ClientErrorCode_ConnectExisted ClientErrorCode = 20004
	ClientErrorCode_PingTimeout    ClientErrorCode = 20005

	ClientErrorCode_SendTimeout ClientErrorCode = 21001

	ClientErrorCode_QueryTimeout ClientErrorCode = 22001
)

var serverClientErrorMap = make(map[int32]ClientErrorCode)

func init() {
	serverClientErrorMap[0] = ClientErrorCode_Success
	serverClientErrorMap[1] = ClientErrorCode_NeedRedirect

	// ConnectAckState_Redirect            int32 = 1
	// ConnectAckState_UnsupportedPlatform int32 = 2
	// ConnectAckState_AuthorizeFailed     int32 = 3
	// ConnectAckState_AuthorizeExpired    int32 = 4
	// ConnectAckState_AppBlock            int32 = 5
	// ConnectAckState_UserBlock           int32 = 6
}
func Trans2ClientErrorCoce(serverCode int32) ClientErrorCode {
	clientCode, ok := serverClientErrorMap[serverCode]
	if ok {
		return clientCode
	} else {
		return ClientErrorCode_Unknown
	}
}
