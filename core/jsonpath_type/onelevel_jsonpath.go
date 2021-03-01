package jsonpath_type

import (
	"encoding/json"
	"github.com/changeJsonStruct/common"
	"github.com/houxiangr/transferOneLevelJson/core"
)

//use open source
//github.com/houxiangr/transferOneLevelJson explain json

type OneLevelJsonpath struct {
	oneTarget map[string]interface{}
}

func (this OneLevelJsonpath) GetValue(expr string) (interface{},error) {
	value,ok := this.oneTarget[expr]
	if !ok {
		return nil,common.JsonPathValueNotExist.SetExtraMsg("err expr is:"+expr)
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

