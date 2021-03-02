package jsonpath_type

import "github.com/houxiangr/changeJsonStruct/common"

const (
	JsonPathTypeOneLevel = "one_level"
	JsonPathTypeExpr     = "expr"
)

type Jsonpath interface {
	GetValue(expr string) (interface{}, error)
	GetType()string
}

func GetJsonPathHandler(jsonPathType string, target string) (Jsonpath, error) {
	switch jsonPathType {
	case JsonPathTypeOneLevel:
		handler := OneLevelJsonpath{}
		err := handler.Init(target)
		if err != nil {
			return nil, err
		}
		return &handler, nil
	case JsonPathTypeExpr:
		handler := ExprJsonpath{}
		err := handler.Init(target)
		if err != nil {
			return nil, err
		}
		return &handler, nil
	default:
		return nil, common.NotHaveJsonPathHandlerType
	}
}
