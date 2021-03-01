package core

import (
	"encoding/json"
	"github.com/houxiangr/changeJsonStruct/common"
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
	"reflect"
)

func ChangeStruct(transferConfStr string, transferTarget string, jsonPathType string) (string, error) {
	transferEntity := make(map[string]interface{})
	err := json.Unmarshal([]byte(transferConfStr), &transferEntity)
	if err != nil {
		return "", err
	}

	jsonPathDeal, err := jsonpath_type.GetJsonPathHandler(jsonPathType, transferTarget)

	resultObj, err := changeStructLogic(transferEntity, jsonPathDeal)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(resultObj)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func changeStructLogic(transferEntity interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	var err error
	switch reflect.TypeOf(transferEntity).Kind() {
	case reflect.Map:
		tempMap := make(map[string]interface{})
		for k, v := range transferEntity.(map[string]interface{}) {
			switch reflect.TypeOf(v).Kind() {
			case reflect.String:
				tempMap[k], err = jsonPathDeal.GetValue(v.(string))
				if err != nil {
					return nil,err
				}
				break
			case reflect.Map:
				vMap := v.(map[string]interface{})
				if common.MapIsHaveKey(vMap, OprKey) {
					tempMap[k], err = dealSpecialOpr(vMap, jsonPathDeal)
					if err != nil {
						return nil, err
					}
					continue
				}
				tempMap[k], err = changeStructLogic(v, jsonPathDeal)
				if err != nil {
					return nil, err
				}
				break
			default:
				tempMap[k], err = changeStructLogic(v, jsonPathDeal)
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
				value, err := jsonPathDeal.GetValue(v.(string))
				if err != nil {
					return nil, err
				}
				tempSlice = append(tempSlice, value)
				break
			default:
				subObj, err := changeStructLogic(v, jsonPathDeal)
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
