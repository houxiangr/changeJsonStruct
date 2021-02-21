package main

import (
	"fmt"
	"github.com/changeJsonStruct/core"
)

func main(){
	transferConf := ``
	transferTarget:=``
	res,err := core.ChangeStruct(transferConf,transferTarget)
	fmt.Println(res)
	fmt.Println(err)
}

