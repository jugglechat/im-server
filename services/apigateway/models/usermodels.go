package models

type UserRegReq struct {
	UserId       string `json:"user_id"`
	Nickname     string `json:"nickname"`
	UserPortrait string `json:"user_portrait"`
}
type UserReqResp struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
