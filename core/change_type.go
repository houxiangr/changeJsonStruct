package core

import (
	"fmt"
	"github.com/changeJsonStruct/common"
	"reflect"
	"strconv"
)

const (
	ToTypeKey = "to_type"
)

const (
	TypeToString = "string"
	TypeToInt    = "int"
)

func changeType(source map[string]interface{}, oneLevelJsonTargetObj map[string]interface{}) (interface{}, error) {
	oprData, ok := source[OprDataKey].(string)
	if !ok {
		return nil, common.OprDataTypeErr
	}
	typeToString, ok := source[ToTypeKey].(string)
	if !ok {
		return nil, common.OprChangeTypeToErr
	}
	targetObj, ok := oneLevelJsonTargetObj[oprData]
	switch reflect.TypeOf(targetObj).Kind() {
	case reflect.Float64:
		switch typeToString {
		case TypeToString:
			return fmt.Sprintf("%+v",targetObj),nil
		case TypeToInt:
			return targetObj,nil
		default:
			return nil, common.OprChangeTypeToErr
		}
	case reflect.String:
		switch typeToString {
		case TypeToString:
			return targetObj,nil
		case TypeToInt:
			return strconv.ParseInt(targetObj.(string),10,64)
		default:
			return nil, common.OprChangeTypeToErr
		}
	default:
		return nil, common.ChangeStructNoSupportType
	}
}
