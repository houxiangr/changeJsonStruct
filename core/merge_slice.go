package core

import (
	"github.com/houxiangr/changeJsonStruct/common"
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"strings"
)

func mergeSlice(source map[string]interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	var err error
	var targetObj interface{}
	oprData := source[OprDataKey].([]interface{})

	resSlice := []interface{}{}
	for _, v := range oprData {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			targetObj, err = jsonPathDeal.GetValue(v.(string))
			//not find target,ignore
			if err != nil && strings.Contains(err.Error(), common.JsonPathValueNotExist.Error()) {
				continue
			}
			if err != nil {
				return nil, err
			}
			targetSlice, ok := targetObj.([]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			resSlice = append(resSlice, targetSlice...)
		case reflect.Slice:
			tempObj := make(map[string]interface{})
			tempObj[OprTypeMergeSlice] = v
			targetObj, err = changeStructLogic(tempObj, jsonPathDeal)
			if err != nil {
				return nil, err
			}
			targetMap, ok := targetObj.(map[string]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			targetSlice, ok := targetMap[OprTypeMergeSlice].([]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			resSlice = append(resSlice, targetSlice...)
		default:
			return nil, common.ChangeStructNoSupportType
		}
	}
	return resSlice, nil
}
