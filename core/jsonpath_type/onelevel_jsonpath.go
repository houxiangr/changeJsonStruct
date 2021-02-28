package jsonpath_type

import (
	"encoding/json"
	"github.com/changeJsonStruct/common"
	"github.com/houxiangr/transferOneLevelJson/core"
)

type OneLevelJsonpath struct {
	oneTarget map[string]interface{}
}

func (this OneLevelJsonpath) GetValue(expr string) (interface{},error) {
	value,ok := this.oneTarget[expr]
	if !ok {
		return nil,common.JsonPathValueNotExist
	}
	return value,nil
}

func (this *OneLevelJsonpath) Init(transferTarget string)error{
	oneLevelJsonTargetObj := make(map[string]interface{})
	oneLevelJsonTarget, err := core.TransferToOneLevelShowAll(transferTarget)
	err = json.Unmarshal([]byte(oneLevelJsonTarget), &oneLevelJsonTargetObj)
	if err != nil {
		return err
	}
	this.oneTarget = oneLevelJsonTargetObj
	return nil
}

