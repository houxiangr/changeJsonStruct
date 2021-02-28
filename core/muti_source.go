package core

import (
	"github.com/changeJsonStruct/common"
	"github.com/changeJsonStruct/core/jsonpath_type"
	"reflect"
)

func mutiSource(source map[string]interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	var targetObj interface{}
	var err error
	oprData := source[OprDataKey].([]interface{})
	for _, v := range oprData {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			targetObj, err = jsonPathDeal.GetValue(v.(string))
			//not find target,ignore
			if err == common.JsonPathValueNotExist {
				continue
			}
			if err != nil {
				return nil, err
			}
			return targetObj, nil
		case reflect.Map:
			targetObj, err := changeStructLogic(v, jsonPathDeal)
			if err != nil {
				return nil, err
			}
			if targetObj != nil {
				return targetObj, nil
			}
		default:
			return nil, common.ChangeStructNoSupportType
		}
	}
	//all not find
	return nil, nil
}
