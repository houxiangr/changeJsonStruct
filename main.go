package main

import (
	"fmt"
	"github.com/houxiangr/changeJsonStruct/core"
	"github.com/houxiangr/changeJsonStruct/core/jsonpath_type"
)

func main() {
	transferConf := `{
	"dup_key1": "$.key1",
	"dup_key2": "$.key2",
	"dup_key3": {
		"dup_key4": "$.key1",
		"dup_key5": "$.key2"
	},
	"dup_key6": [
		"$.key1",
		{
			"dup_key7": "$.key1"
		}
	],
	"dup_key7": {
		"opr": "merge_map",
		"data": [
			"$.key3",
			{
				"dup_key8": "$.key1"
			}
		]
	},
	"dup_key9": {
		"opr": "muti_source",
		"data": [
			"$.key5",
			"$.key2"
		]
	},
	"dup_key10": {
		"opr": "default_value",
		"data": "10"
	},
	"dup_key11": {
		"opr": "merge_slice",
		"data": [
			"$.key5", [
				"$.key1",
				"$.key2",
				{
					"dup_key12": "$.key1"
				}
			]
		]
	},
	"key8": {
		"opr": "change_type",
		"data": "$.key1",
		"to_type": "string"
	}
}`
	transferTarget := `{
	"key1": 1,
	"key2": 2,
	"key3": {
		"key4": 4
	},
	"key5": [
		1,
		2,
		3
	]
}`
	res, err := core.ChangeStruct(transferConf, transferTarget, jsonpath_type.JsonPathTypeOneLevel)
	fmt.Println(res)
	fmt.Println(err)
}
