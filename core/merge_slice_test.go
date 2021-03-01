package core

import (
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"testing"
)

var mergeSliceJsonPathDeal jsonpath_type.Jsonpath

func init() {
	target := `{
	"key1":1,
	"key3":[1,2], 
	"key5":[3,4]
}`

	mergeSliceJsonPathDeal, _ = jsonpath_type.GetJsonPathHandler(jsonpath_type.JsonPathTypeOneLevel, target)
}

func Test_mergeSlice(t *testing.T) {
	type args struct {
		source map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wanterr error
	}{
		{
			name: "merge two exist slice",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key3",
						"$.key5",
					},
				},
			},
			want:    []interface{}{float64(1), float64(2), float64(3), float64(4)},
			wanterr: nil,
		},
		{
			name: "merge one exist slice and one create slice",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key3",
						[]interface{}{
							"$.key1",
						},
					},
				},
			},
			want:    []interface{}{float64(1), float64(2), float64(1)},
			wanterr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeSlice(tt.args.source, mergeSliceJsonPathDeal)
			if err != tt.wanterr {
				t.Errorf("mergeSlice(),err=%+v ", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
