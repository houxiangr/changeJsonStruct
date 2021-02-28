package core

import (
	"github.com/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"testing"
)

var changeTypeJsonPathDeal jsonpath_type.Jsonpath

func init() {
	target := `{
	"key1": 1,
	"key2": "2",
	"key3": {
		"key4": 4
	}
}`

	changeTypeJsonPathDeal, _ = jsonpath_type.GetJsonPathHandler(jsonpath_type.JsonPathTypeOneLevel, target)
}

func Test_changeType(t *testing.T) {
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
			name: "int to string",
			args: args{
				source: map[string]interface{}{
					OprDataKey: "$.key1",
					TypeToKey:  TypeToString,
				},
			},
			want:    "1",
			wanterr: nil,
		},
		{
			name: "string to int",
			args: args{
				source: map[string]interface{}{
					OprDataKey: "$.key2",
					TypeToKey:  TypeToInt,
				},
			},
			want:    int64(2),
			wanterr: nil,
		},
		{
			name: "struct to string",
			args: args{
				source: map[string]interface{}{
					OprDataKey: "$.key3",
					TypeToKey:  TypeToString,
				},
			},
			want:    `{"key4":4}`,
			wanterr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := changeType(tt.args.source, changeTypeJsonPathDeal)
			if err != tt.wanterr {
				t.Errorf("changeType(),err=%+v ", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeType() = %v, want %v", got, tt.want)
			}
		})
	}
}
