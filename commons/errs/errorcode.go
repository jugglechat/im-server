package errs

/*
0 : success
10000~10999 : api
11000~11999 : connect
12000~12999 : private msg
13000~13999 : group
14000~14999 : chatroom

*/

type ErrorMsg struct {
	HttpCode int32  `json:"-"`
	Code     int32  `json:"code"`
	Msg      string `json:"msg"`
}

func NewErrorMsg(httpCode, code int32, msg string) *ErrorMsg {
	return &ErrorMsg{
		Code: code,
		Msg:  msg,
	}
}

var ERR_SUCCESS *ErrorMsg = NewErrorMsg(200, 0, "success")

//api
var ERR_API_DEFAULT = NewErrorMsg(200, 10000, "default error.")
var ERR_API_APPKEY_REQUIRED = NewErrorMsg(400, 10001, "appkey is required")
var ERR_API_NONCE_REQUIRED = NewErrorMsg(400, 10002, "nonce is required")
var ERR_API_TIMESTAMP_REQUIRED = NewErrorMsg(400, 1003, "timestamp is required")
var ERR_API_SIGNATURE_REQUIRED = NewErrorMsg(400, 10004, "signature is required")
var ERR_API_APP_NOT_EXISTED = NewErrorMsg(500, 10005, "app not existed")
var ERR_API_SIGNATURE_FAIL = NewErrorMsg(403, 10006, "signature is fail")
var ERR_API_REQ_BODY_ILLEGAL = NewErrorMsg(400, 10007, "request body is illegal")
var ERR_API_INTERNAL_TIMEOUT = NewErrorMsg(500, 10008, "internal service timeout")
var ERR_API_INTERNAL_RESP_FAIL = NewErrorMsg(500, 10009, "internal service error")

//conn
var ERR_CONN_DEFAULT = NewErrorMsg(200, 11000, "default error.")
var ERR_CONN_APPKEY_REQUIRED = NewErrorMsg(400, 11001, "")
var ERR_CONN_TOKEN_REQUIRED = NewErrorMsg(400, 11002, "")
var ERR_CONN_APP_NOT_EXISTED = NewErrorMsg(500, 11003, "")
var ERR_CONN_TOKEN_ILLEGAL = NewErrorMsg(403, 11004, "")
var ERR_CONN_TOKEN_AUTHFAIL = NewErrorMsg(401, 11005, "")
var ERR_CONN_TOKEN_EXPIRED = NewErrorMsg(401, 11006, "")

/*
type ErrorCode int32
const (
	ERR_SUCCESS ErrorCode = 0
	//api
	ERR_API_DEFAULT ErrorCode = 10000
	//conn
	ERR_CONN_DEFAULT         ErrorCode = 11000
	ERR_CONN_APPKEY_REQUIRED ErrorCode = 11001
	ERR_CONN_TOKEN_REQUIRED  ErrorCode = 11002
	ERR_CONN_APP_NOT_EXISTED ErrorCode = 11003
	ERR_CONN_TOKEN_ILLEGAL   ErrorCode = 11004
	ERR_CONN_TOKEN_AUTHFAIL  ErrorCode = 11005
	ERR_CONN_TOKEN_EXPIRED   ErrorCode = 11006
)
*/
