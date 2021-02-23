package core

import (
	"encoding/json"
	"github.com/changeJsonStruct/common"
	"github.com/houxiangr/transferOneLevelJson/core"
	"reflect"
)

func ChangeStruct(transferConfStr string, transferTarget string) (string, error) {
	transferEntity := make(map[string]interface{})
	err := json.Unmarshal([]byte(transferConfStr), &transferEntity)
	if err != nil {
		return "", err
	}

	oneLevelJsonTargetObj := make(map[string]interface{})
	oneLevelJsonTarget, err := core.TransferToOneLevelShowAll(transferTarget)
	err = json.Unmarshal([]byte(oneLevelJsonTarget), &oneLevelJsonTargetObj)
	if err != nil {
		return "", err
	}

	resultObj, err := changeStructLogic(transferEntity, oneLevelJsonTargetObj)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(resultObj)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func changeStructLogic(transferEntity interface{}, oneLevelJsonTargetObj map[string]interface{}) (map[string]interface{}, error){
	var err error
	switch reflect.TypeOf(transferEntity).Kind() {
	case reflect.Map:
		tempMap := make(map[string]interface{})
		for k,v := range transferEntity.(map[string]interface{}) {
			switch reflect.TypeOf(v).Kind(){
			case reflect.String:
				tempMap[k] = oneLevelJsonTargetObj[v.(string)]
			default:
				tempMap[k],err = changeStructLogic(v,oneLevelJsonTargetObj)
				if err != nil {
					return nil,err
				}
			}
		}
		return tempMap,nil
	case reflect.Slice:
	default:
		return nil, common.ChangeStructNoSupportType
	}
	return nil, nil
}