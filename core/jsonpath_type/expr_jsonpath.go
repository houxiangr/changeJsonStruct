package jsonpath_type

import (
	"encoding/json"
	"github.com/changeJsonStruct/common"
	"github.com/oliveagle/jsonpath"
	"strings"
)

//use open source
//github.com/oliveagle/jsonpath explain json

type ExprJsonpath struct {
	target interface{}
}

func (this ExprJsonpath) GetValue(expr string) (interface{}, error) {
	res, err := jsonpath.JsonPathLookup(this.target, expr)
	if err != nil && strings.Contains(err.Error(), "not found in object") {
		return nil, common.JsonPathValueNotExist.SetExtraMsg("err expr is:" + expr)
	}
	return res, err
}

func (this *ExprJsonpath) Init(transferTarget string) error {
	var jsonTargetObj interface{}
	err := json.Unmarshal([]byte(transferTarget), &jsonTargetObj)
	if err != nil {
		return err
	}
	this.target = jsonTargetObj
	return nil
}
