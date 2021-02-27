package core

import (
	"encoding/json"
	"github.com/changeJsonStruct/common"
	"github.com/houxiangr/transferOneLevelJson/core"
	"reflect"
)

const (
	OprKey     = "opr"
	OprDataKey = "data"
)

const (
	OprTypeDefaultValue = "default_value"
	OprTypeMergeMap     = "merge_map"
	OprTypeMergeSlice   = "merge_slice"
	OprTypeMutiSource   = "muti_source"
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

func changeStructLogic(transferEntity interface{}, oneLevelJsonTargetObj map[string]interface{}) (interface{}, error) {
	var err error
	switch reflect.TypeOf(transferEntity).Kind() {
	case reflect.Map:
		tempMap := make(map[string]interface{})
		for k, v := range transferEntity.(map[string]interface{}) {
			switch reflect.TypeOf(v).Kind() {
			case reflect.String:
				tempMap[k] = oneLevelJsonTargetObj[v.(string)]
				break
			case reflect.Map:
				vMap := v.(map[string]interface{})
				if common.MapIsHaveKey(vMap, OprKey) {
					tempMap[k], err = dealSpecialOpr(vMap, oneLevelJsonTargetObj)
					if err != nil {
						return nil, err
					}
					continue
				}
				tempMap[k], err = changeStructLogic(v, oneLevelJsonTargetObj)
				if err != nil {
					return nil, err
				}
				break
			default:
				tempMap[k], err = changeStructLogic(v, oneLevelJsonTargetObj)
				if err != nil {
					return nil, err
				}
				break
			}
		}
		return tempMap, nil
	case reflect.Slice:
		tempSlice := []interface{}{}
		for _, v := range transferEntity.([]interface{}) {
			switch reflect.TypeOf(v).Kind() {
			case reflect.String:
				tempSlice = append(tempSlice, oneLevelJsonTargetObj[v.(string)])
				break
			default:
				subObj, err := changeStructLogic(v, oneLevelJsonTargetObj)
				if err != nil {
					return nil, err
				}
				tempSlice = append(tempSlice, subObj)
				break
			}
		}
		return tempSlice, nil
	default:
		return nil, common.ChangeStructNoSupportType
	}
}

func dealSpecialOpr(source map[string]interface{}, oneLevelJsonTargetObj map[string]interface{}) (interface{}, error) {
	switch source[OprKey].(string) {
	case OprTypeMergeMap:
		return mergeMap(source, oneLevelJsonTargetObj)
	case OprTypeMergeSlice:
		return mergeSlice(source, oneLevelJsonTargetObj)
	case OprTypeMutiSource:
		return mutiSource(source, oneLevelJsonTargetObj)
	case OprTypeDefaultValue:
		return defaultValue(source)
	default:
		return nil, common.ChangeStructNoSupportOpr
	}
}

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

func defaultValue(source map[string]interface{}) (interface{}, error) {
	return source[OprDataKey], nil
}

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
