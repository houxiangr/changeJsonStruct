package jsonpath_type

import (
	"github.com/changeJsonStruct/common"
	"reflect"
	"testing"
)

func TestGetJsonPathHandler(t *testing.T) {
	type args struct {
		jsonPathType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wanterr error
	}{
		{
			name: "one level jsonpath handler",
			args: args{
				jsonPathType: JsonPathTypeOneLevel,
			},
			want:    "*jsonpath_type.OneLevelJsonpath",
			wanterr: nil,
		},
		{
			name: "not match jsonpath handler",
			args: args{
				jsonPathType: "not_exist_type",
			},
			want:    "",
			wanterr: common.NotHaveJsonPathHandlerType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJsonPathHandler(tt.args.jsonPathType, "{}")
			if err != tt.wanterr {
				t.Errorf("ChangeStruct(),err=%+v ", err)
			}
			if got != nil && !reflect.DeepEqual(reflect.TypeOf(got).String(), tt.want) {
				t.Errorf("ChangeStruct() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
			}
		})
	}
}
