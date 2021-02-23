package common

import "fmt"

type Error struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
}

func NewError(err Error, msg string) Error {
	return Error{err.ErrNo, msg}
}

func (e Error) Error() string {
	return e.ErrMsg
}

func (e Error) GetErrNo() int {
	return e.ErrNo
}

func (e Error) SetExtraMsg(s string) Error {
	e.ErrMsg = fmt.Sprintf("%s:%s", e.ErrMsg, s)
	return e
}

var (
	Success = Error{0, "success"}
	//基础错误1000~1999
	ChangeStructNoSupportType            = Error{1001, "not support type"}
)
