package core

import (
	"encoding/json"
	"fmt"
	"github.com/changeJsonStruct/common"
	"github.com/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"strconv"
)

const (
	TypeToKey = "to_type"
)

const (
	TypeToString = "string"
	TypeToInt    = "int"
)

func changeType(source map[string]interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	oprData, ok := source[OprDataKey].(string)
	if !ok {
		return nil, common.OprDataTypeErr
	}
	typeTo, ok := source[TypeToKey].(string)
	if !ok {
		return nil, common.OprChangeTypeToErr
	}
	targetObj, _ :=  jsonPathDeal.GetValue(oprData)
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
