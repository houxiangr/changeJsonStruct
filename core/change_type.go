package core

import (
	"encoding/json"
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
	typeTo, ok := source[ToTypeKey].(string)
	if !ok {
		return nil, common.OprChangeTypeToErr
	}
	targetObj, ok := oneLevelJsonTargetObj[oprData]
	switch reflect.TypeOf(targetObj).Kind() {
	case reflect.Float64:
		switch typeTo {
		case TypeToString:
			return fmt.Sprintf("%+v",targetObj),nil
		case TypeToInt:
			return targetObj,nil
		default:
			return nil, common.OprChangeTypeToErr
		}
	case reflect.String:
		switch typeTo {
		case TypeToString:
			return targetObj,nil
		case TypeToInt:
			return strconv.ParseInt(targetObj.(string),10,64)
		default:
			return nil, common.OprChangeTypeToErr
		}
	case reflect.Map,reflect.Slice:
		switch typeTo{
		case TypeToString:
			v,err := json.Marshal(targetObj)
			if err != nil {
				return nil,err
			}
			return string(v),nil
		default:
			return nil, common.OprChangeTypeToErr
		}

	default:
		return nil, common.ChangeStructNoSupportType
	}
}
