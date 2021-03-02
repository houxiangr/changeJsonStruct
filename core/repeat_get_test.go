package core

import (
	"github.com/houxiangr/changeJsonStruct/common"
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"strings"
	"testing"
)

var repeatGetJsonPathDeal jsonpath_type.Jsonpath

func init() {
	target := `{
	"key1":"{\"key2\":2}"
}`

	repeatGetJsonPathDeal, _ = jsonpath_type.GetJsonPathHandler(jsonpath_type.JsonPathTypeOneLevel, target)
}

func Test_repeat_get(t *testing.T) {
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
			name: "repeat get",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key1",
						"$.key2",
					},
				},
			},
			want:    float64(2),
			wanterr: nil,
		},
		{
			name: "repeat get",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key1",
						map[string]interface{}{
							"map_key": "$.key2",
						},
					},
				},
			},
			want: map[string]interface{}{
				"map_key": float64(2),
			},
			wanterr: nil,
		},
		{
			name: "repeat get",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key1",
						[]interface{}{
							"$.key2",
						},
					},
				},
			},
			want: []interface{}{
				float64(2),
			},
			wanterr: nil,
		},
		{
			name: "repeat get",
			args: args{
				source: map[string]interface{}{
					OprDataKey: []interface{}{
						"$.key1",
						[]interface{}{
							"$.key2",
						},
						"$.not_support",
					},
				},
			},
			want:    nil,
			wanterr: common.RepeatGetMiddleDataTypeErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repeatGet(tt.args.source, repeatGetJsonPathDeal)
			if err != nil && !strings.Contains(err.Error(), tt.wanterr.Error()) {
				t.Errorf("repeatGet() err = %v, want %v", err.Error(), tt.wanterr.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repeatGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
