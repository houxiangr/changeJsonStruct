package core

import "github.com/changeJsonStruct/common"

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
