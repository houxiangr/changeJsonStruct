package core

import (
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"testing"
)

var mergeMapJsonPathDeal jsonpath_type.Jsonpath

func init() {
	target := `{
	"key1":1,
	"key3": {
		"key4": 4
	},
	"key5":{
		"key6": 6
	}
}`

	mergeMapJsonPathDeal, _ = jsonpath_type.GetJsonPathHandler(jsonpath_type.JsonPathTypeOneLevel, target)
}

func Test_mergeMap(t *testing.T) {
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
			name: "merge two exist map",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key3",
						"$.key5",
					},
				},
			},
			want: map[string]interface{}{
				"key4": float64(4),
				"key6": float64(6),
			},
			wanterr: nil,
		},
		{
			name: "merge one exist map and one create map",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key3",
						map[string]interface{}{
							"key7": "$.key1",
						},
					},
				},
			},
			want: map[string]interface{}{
				"key4": float64(4),
				"key7": float64(1),
			},
			wanterr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeMap(tt.args.source, mergeMapJsonPathDeal)
			if err != tt.wanterr {
				t.Errorf("mergeMap(),err=%+v ", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
