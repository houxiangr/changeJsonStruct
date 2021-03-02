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
	ChangeStructNoSupportType  = Error{1001, "not support type"}
	ChangeStructNoSupportOpr   = Error{1002, "not support opr"}
	OprDataTypeErr             = Error{1003, "opr data type err"}
	OprChangeTypeToErr         = Error{1004, "change type to err"}
	NotHaveJsonPathHandlerType = Error{1005, "not have json path handler type"}
	JsonPathValueNotExist      = Error{1006, "json path value not exist"}
	RepeatGetMiddleDataTypeErr = Error{1007,"repeat get data middle data err"}
)
