package clients

import "fmt"

type ConnectTimeoutErr struct {
	Err error
}

type ConnectFailedErr struct {
	Err error
}

func (e ConnectTimeoutErr) Error() string {
	return fmt.Sprintf("connect timeout %v", e.Err)
}

func (e ConnectFailedErr) Error() string {
	return fmt.Sprintf("connect failed %v", e.Err)
}
