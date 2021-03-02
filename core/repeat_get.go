package core

import (
	"encoding/json"
	"fmt"
	"github.com/houxiangr/changeJsonStruct/common"
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
	"reflect"
)

func repeatGet(source map[string]interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	var err error
	var nextTargetStr interface{}
	var nextTargetObj interface{}
	oprData := source[OprDataKey].([]interface{})

	if len(oprData) == 0 {
		return nil, nil
	}
	//repeat get first opr must be string
	firstOpr, ok := oprData[0].(string)
	if !ok {
		return nil, common.OprDataTypeErr
	}
	nextTargetObj, err = jsonPathDeal.GetValue(firstOpr)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(oprData); i++ {
		nextTargetStr = ""
		switch reflect.TypeOf(nextTargetObj).Kind() {
		case reflect.String:
			nextTargetStr = nextTargetObj.(string)
		case reflect.Map:
			nextBytes, err := json.Marshal(nextTargetObj)
			if err != nil {
				return nil, err
			}
			nextTargetStr = string(nextBytes)
		default:
			return nil, common.RepeatGetMiddleDataTypeErr.SetExtraMsg(fmt.Sprintf("%+v",oprData[i]))
		}

		switch reflect.TypeOf(oprData[i]).Kind() {
		case reflect.String:
			transferConfObj := make(map[string]interface{})
			transferConfObj[OprTypeRepeatGet+"string"] = oprData[i]
			transferConfBytes, err := json.Marshal(transferConfObj)
			if err != nil {
				return nil, err
			}
			targetStr, err := ChangeStruct(string(transferConfBytes), nextTargetStr.(string), jsonPathDeal.GetType())
			if err != nil {
				return nil, err
			}
			var targetObj map[string]interface{}
			err = json.Unmarshal([]byte(targetStr), &targetObj)
			if err != nil {
				return nil, err
			}
			nextTargetObj = targetObj[OprTypeRepeatGet+"string"]
		case reflect.Map:
			transferConfBytes, err := json.Marshal(oprData[i])
			if err != nil {
				return nil, err
			}
			nextTargetStr, err = ChangeStruct(string(transferConfBytes), nextTargetStr.(string), jsonPathDeal.GetType())
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(nextTargetStr.(string)), &nextTargetObj)
			if err != nil {
				return nil, err
			}
		case reflect.Slice:
			transferConfObj := make(map[string]interface{})
			transferConfObj[OprTypeRepeatGet+"slice"] = oprData[i]
			transferConfBytes, err := json.Marshal(transferConfObj)
			if err != nil {
				return nil, err
			}
			targetStr, err := ChangeStruct(string(transferConfBytes), nextTargetStr.(string), jsonPathDeal.GetType())
			if err != nil {
				return nil, err
			}
			var targetObj map[string]interface{}
			err = json.Unmarshal([]byte(targetStr), &targetObj)
			if err != nil {
				return nil, err
			}
			nextTargetObj = targetObj[OprTypeRepeatGet+"slice"]
		default:
			return nil, common.ChangeStructNoSupportType
		}
	}
	//all not find
	return nextTargetObj, nil
}
