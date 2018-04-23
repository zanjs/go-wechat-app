package config

const (
	Success = iota
	ErrorAuth
	ErrorGetUser
	ErrorUserRecord
	ErrorValidate
	ErrorDayScan
	ErrorTotalScan
	ErrorWxValidate
	ErrorWxLogin
	ErrorWxUserInfo
)

var ErrMsg map[int]string

func init() {
	ErrMsg = make(map[int]string)
	ErrMsg[ErrorAuth] = "auth failure"
	ErrMsg[ErrorGetUser] = "get user error"
	ErrMsg[ErrorUserRecord] = "get user records error"
	ErrMsg[ErrorValidate] = "validate store data failure"
	ErrMsg[ErrorDayScan] = "scan the max times of day"
	ErrMsg[ErrorTotalScan] = "only can scan 500 books"
	ErrMsg[ErrorWxValidate] = "validate login data failure"
	ErrMsg[ErrorWxLogin] = "wx login failed"
	ErrMsg[ErrorWxUserInfo] = "get user_info failed"
}
