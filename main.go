package main

import (
	"fmt"
	"github.com/changeJsonStruct/core"
)

func main(){
	transferConf := `{
	"dup_key1": "$.key1",
	"dup_key2": "$.key2"
}`
	transferTarget:=`{
    "key1":1,
    "key2":2
}`
	res,err := core.ChangeStruct(transferConf,transferTarget)
	fmt.Println(res)
	fmt.Println(err)
}

