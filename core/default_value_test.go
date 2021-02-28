package core

import (
	"reflect"
	"testing"
)

func Test_defaultValue(t *testing.T) {
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
			name: "default value",
			args: args{
				source: map[string]interface{}{
					OprDataKey: 1,
				},
			},
			want:    1,
			wanterr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultValue(tt.args.source)
			if err != tt.wanterr {
				t.Errorf("changeType(),err=%+v ", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeType() = %v, want %v", got, tt.want)
			}
		})
	}
}
