package core

import (
	"github.com/changeJsonStruct/common"
	"github.com/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"strings"
)

func mergeMap(source map[string]interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	var err error
	var targetObj interface{}
	oprData := source[OprDataKey].([]interface{})

	resMap := make(map[string]interface{})
	for _, v := range oprData {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			targetObj, err = jsonPathDeal.GetValue(v.(string))
			//not find target,ignore
			if err != nil && strings.Contains(err.Error(), common.JsonPathValueNotExist.Error()) {
				continue
			}
			if err != nil {
				return nil,err
			}
			targetMap, ok := targetObj.(map[string]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			common.MergeMap(resMap, targetMap)
		case reflect.Map:
			targetObj, err = changeStructLogic(v, jsonPathDeal)
			if err != nil {
				return nil, err
			}
			targetMap, ok := targetObj.(map[string]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			common.MergeMap(resMap, targetMap)
		default:
			return nil, common.ChangeStructNoSupportType
		}
	}
	return resMap, nil
}

