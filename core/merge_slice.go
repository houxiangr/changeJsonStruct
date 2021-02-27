package core

import (
	"github.com/changeJsonStruct/common"
	"reflect"
)

func mergeSlice(source map[string]interface{}, oneLevelJsonTargetObj map[string]interface{}) (interface{}, error) {
	oprData := source[OprDataKey].([]interface{})

	resSlice := []interface{}{}
	for _, v := range oprData {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			targetObj, ok := oneLevelJsonTargetObj[v.(string)]
			//not find target,ignore
			if !ok {
				continue
			}
			targetSlice, ok := targetObj.([]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			resSlice = append(resSlice, targetSlice...)
		case reflect.Slice:
			tempObj := make(map[string]interface{})
			tempObj[OprTypeMergeSlice] = v
			targetObj, err := changeStructLogic(tempObj, oneLevelJsonTargetObj)
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
