package jsonpath_type

import "github.com/changeJsonStruct/common"

const (
	JsonPathTypeOneLevel = "one_level"
)

type Jsonpath interface {
	GetValue(expr string) (interface{}, error)
}


func GetJsonPathHandler(jsonPathType string,target string)(Jsonpath,error){
	switch jsonPathType {
	case JsonPathTypeOneLevel:
		handler := OneLevelJsonpath{}
		err := handler.Init(target)
		if err != nil {
			return nil,err
		}
		return &handler,nil
	default:
		return nil,common.NotHaveJsonPathHandlerType
	}
}