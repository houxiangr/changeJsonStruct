package core

import (
	"github.com/changeJsonStruct/common"
	"reflect"
)

func mutiSource(source map[string]interface{}, oneLevelJsonTargetObj map[string]interface{}) (interface{}, error) {
	oprData := source[OprDataKey].([]interface{})
	for _, v := range oprData {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			targetObj, ok := oneLevelJsonTargetObj[v.(string)]
			//not find target,ignore
			if !ok {
				continue
			} else {
				return targetObj, nil
			}

		case reflect.Map:
			targetObj, err := changeStructLogic(v, oneLevelJsonTargetObj)
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
