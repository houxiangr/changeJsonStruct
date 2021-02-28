package core

import (
	"github.com/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"testing"
)

var mutiSourceJsonPathDeal jsonpath_type.Jsonpath

func init() {
	target := `{
	"key1":1,
	"key3":[1,2], 
	"key5":[3,4]
}`

	mutiSourceJsonPathDeal, _ = jsonpath_type.GetJsonPathHandler(jsonpath_type.JsonPathTypeOneLevel, target)
}

func Test_muti_source(t *testing.T) {
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
			name: "first hit",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key3",
						"$.key5",
					},
				},
			},
			want:    []interface{}{float64(1), float64(2)},
			wanterr: nil,
		},
		{
			name: "second hit",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key_not_exist",
						"$.key3",
					},
				},
			},
			want:    []interface{}{float64(1), float64(2)},
			wanterr: nil,
		},
		{
			name: "second hit create map",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key_not_exist",
						map[string]interface{}{
							"key6": "$.key1",
						},
					},
				},
			},
			want: map[string]interface{}{
				"key6": float64(1),
			},
			wanterr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mutiSource(tt.args.source, mutiSourceJsonPathDeal)
			if err != tt.wanterr {
				t.Errorf("mutiSource(),err=%+v ", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutiSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
