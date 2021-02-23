package main

import (
	"fmt"
	"github.com/changeJsonStruct/core"
)

func main(){
	transferConf := `{
	"dup_key1": "$.key1",
	"dup_key2": "$.key2",
	"dup_key3": {
		"dup_key4":"$.key1",
		"dup_key5":"$.key2"
	},
	"dup_key6":[
		"$.key1",
		{
			"dup_key7":"$.key1"
		}
	]
}`
	transferTarget:=`{
    "key1":1,
    "key2":2
}`
	res,err := core.ChangeStruct(transferConf,transferTarget)
	fmt.Println(res)
	fmt.Println(err)
}

