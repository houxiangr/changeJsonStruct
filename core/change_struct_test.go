package core

import (
	"github.com/changeJsonStruct/common"
	"github.com/changeJsonStruct/core/jsonpath_type"
	"reflect"
	"strings"
	"testing"
)

func TestChangeStruct(t *testing.T) {
	type args struct {
		transferConfStr string
		transferTarget  string
		jsonPathType    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wanterr error
	}{
		{
			name: "one level normal process",
			args: args{
				transferConfStr: `{"dup_key1":"$.key1","dup_key2":"$.key2","dup_key3":{"dup_key4":"$.key1","dup_key5":"$.key2"},"dup_key6":["$.key1",{"dup_key7":"$.key1"}],"dup_key7":{"opr":"merge_map","data":["$.key3",{"dup_key8":"$.key1"}]},"dup_key9":{"opr":"muti_source","data":["$.key5","$.key2"]},"dup_key10":{"opr":"default_value","data":"10"},"dup_key11":{"opr":"merge_slice","data":["$.key5",["$.key1","$.key2",{"dup_key12":"$.key1"}]]},"key8":{"opr":"change_type","data":"$.key1","to_type":"string"}}`,
				transferTarget:  `{"key1":1,"key2":2,"key3":{"key4":4},"key5":[1,2,3]}`,
				jsonPathType:    jsonpath_type.JsonPathTypeOneLevel,
			},
			want:    `{"dup_key1":1,"dup_key10":"10","dup_key11":[1,2,3,1,2,{"dup_key12":1}],"dup_key2":2,"dup_key3":{"dup_key4":1,"dup_key5":2},"dup_key6":[1,{"dup_key7":1}],"dup_key7":{"dup_key8":1,"key4":4},"dup_key9":[1,2,3],"key8":"1"}`,
			wanterr: nil,
		},
		{
			name: "one level empty transferConfStr",
			args: args{
				transferConfStr: `{}`,
				transferTarget:  `{"key1":1,"key2":2,"key3":{"key4":4},"key5":[1,2,3]}`,
				jsonPathType:    jsonpath_type.JsonPathTypeOneLevel,
			},
			want:    `{}`,
			wanterr: nil,
		},
		{
			name: "one level empty transferTarget",
			args: args{
				transferConfStr: `{"dup_key1":"$.key1"}`,
				transferTarget:  `{}`,
				jsonPathType:    jsonpath_type.JsonPathTypeOneLevel,
			},
			want:    ``,
			wanterr: common.JsonPathValueNotExist,
		},
		{
			name: "expr normal process",
			args: args{
				transferConfStr: `{"dup_key1":"$.key1","dup_key2":"$.key2","dup_key3":{"dup_key4":"$.key1","dup_key5":"$.key2"},"dup_key6":["$.key1",{"dup_key7":"$.key1"}],"dup_key7":{"opr":"merge_map","data":["$.key3",{"dup_key8":"$.key1"}]},"dup_key9":{"opr":"muti_source","data":["$.key5","$.key2"]},"dup_key10":{"opr":"default_value","data":"10"},"dup_key11":{"opr":"merge_slice","data":["$.key5",["$.key1","$.key2",{"dup_key12":"$.key1"}]]},"key8":{"opr":"change_type","data":"$.key1","to_type":"string"}}`,
				transferTarget:  `{"key1":1,"key2":2,"key3":{"key4":4},"key5":[1,2,3]}`,
				jsonPathType:    jsonpath_type.JsonPathTypeExpr,
			},
			want:    `{"dup_key1":1,"dup_key10":"10","dup_key11":[1,2,3,1,2,{"dup_key12":1}],"dup_key2":2,"dup_key3":{"dup_key4":1,"dup_key5":2},"dup_key6":[1,{"dup_key7":1}],"dup_key7":{"dup_key8":1,"key4":4},"dup_key9":[1,2,3],"key8":"1"}`,
			wanterr: nil,
		},
		{
			name: "expr empty transferConfStr",
			args: args{
				transferConfStr: `{}`,
				transferTarget:  `{"key1":1,"key2":2,"key3":{"key4":4},"key5":[1,2,3]}`,
				jsonPathType:    jsonpath_type.JsonPathTypeOneLevel,
			},
			want:    `{}`,
			wanterr: nil,
		},
		{
			name: "expr empty transferTarget",
			args: args{
				transferConfStr: `{"dup_key1":"$.key1"}`,
				transferTarget:  `{}`,
				jsonPathType:    jsonpath_type.JsonPathTypeExpr,
			},
			want:    ``,
			wanterr: common.JsonPathValueNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChangeStruct(tt.args.transferConfStr, tt.args.transferTarget, tt.args.jsonPathType)
			if err != nil && tt.wanterr != nil && !strings.Contains(err.Error(), tt.wanterr.Error()) {
				t.Errorf("ChangeStruct(),err=%+v, wanterr %+v ", err.Error(), tt.wanterr.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
