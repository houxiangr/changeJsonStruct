package core

import (
	"github.com/houxiangr/changeJsonStruct/common"
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
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
	OprTypeChangeType   = "change_type"
	OprTypeRepeatGet    = "repeat_get"
)

func dealSpecialOpr(source map[string]interface{}, jsonPathDeal jsonpath_type.Jsonpath) (interface{}, error) {
	switch source[OprKey].(string) {
	case OprTypeMergeMap:
		return mergeMap(source, jsonPathDeal)
	case OprTypeMergeSlice:
		return mergeSlice(source, jsonPathDeal)
	case OprTypeMutiSource:
		return mutiSource(source, jsonPathDeal)
	case OprTypeDefaultValue:
		return defaultValue(source)
	case OprTypeChangeType:
		return changeType(source, jsonPathDeal)
	case OprTypeRepeatGet:
		return repeatGet(source, jsonPathDeal)
	default:
		return nil, common.ChangeStructNoSupportOpr
	}
}
