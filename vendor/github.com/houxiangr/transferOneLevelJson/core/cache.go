package core

import "sync"

var cacheMap map[string]string
var cacheLock sync.RWMutex

func init(){
	cacheMap = make(map[string]string)
	cacheLock = sync.RWMutex{}
}
