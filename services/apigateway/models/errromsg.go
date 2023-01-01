package models

type ErrorMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func DefaultErr() *ErrorMsg {
	return &ErrorMsg{
		Code: 10000,
		Msg:  "Default Error Msg.",
	}
}
