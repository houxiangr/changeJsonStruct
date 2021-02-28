package common

import (
	"reflect"
	"testing"
)

func TestMergeMap(t *testing.T) {
	type args struct {
		map1 map[string]interface{}
		map2 map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "merge map,not cover",
			args: args{
				map1: map[string]interface{}{
					"key1": 1,
				},
				map2: map[string]interface{}{
					"key2": 2,
				},
			},
			want: map[string]interface{}{
				"key1": 1,
				"key2": 2,
			},
		},
		{
			name: "merge map,cover",
			args: args{
				map1: map[string]interface{}{
					"key1": 1,
				},
				map2: map[string]interface{}{
					"key1": 2,
				},
			},
			want: map[string]interface{}{
				"key1": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeMap(tt.args.map1, tt.args.map2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIsHaveKey(t *testing.T) {
	type args struct {
		map1 map[string]interface{}
		key  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "merge map,not cover",
			args: args{
				map1: map[string]interface{}{
					"key1": 1,
				},
				key: "key1",
			},
			want: true,
		},
		{
			name: "merge map,cover",
			args: args{
				map1: map[string]interface{}{
					"key1": 1,
				},
				key: "key2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapIsHaveKey(tt.args.map1, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapIsHaveKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
