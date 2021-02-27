package core

import (
	"github.com/changeJsonStruct/common"
	"reflect"
)

func mergeMap(source map[string]interface{}, oneLevelJsonTargetObj map[string]interface{}) (interface{}, error) {
	oprData := source[OprDataKey].([]interface{})

	resMap := make(map[string]interface{})
	for _, v := range oprData {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			targetObj, ok := oneLevelJsonTargetObj[v.(string)]
			//not find target,ignore
			if !ok {
				continue
			}
			targetMap, ok := targetObj.(map[string]interface{})
			if !ok {
				return nil, common.OprDataTypeErr
			}
			common.MergeMap(resMap, targetMap)
		case reflect.Map:
			targetObj, err := changeStructLogic(v, oneLevelJsonTargetObj)
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

