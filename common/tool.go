package common

func MapIsHaveKey(sourceMap interface{},key string) bool {
	if _,ok := sourceMap.(map[string]interface{})[key];ok{
		return true
	}
	return false
}

func MergeMap(map1 map[string]interface{},map2 map[string]interface{}) map[string]interface{} {
	for k,v := range map2 {
		tempv := v
		map1[k] = tempv
	}
	return map1
}
