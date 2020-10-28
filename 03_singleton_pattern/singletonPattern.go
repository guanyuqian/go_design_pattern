package _3_singleton_pattern

import "sync"

//步骤 1
//创建一个 Singleton 类。
var lazyInstance *Instance
var hungryInstance *Instance

func GetHungryInstance() *Instance {
	return hungryInstance
}

func GetLazyInstance() *Instance {
	once.Do(func() {
		lazyInstance = new(Instance)
		lazyInstance.Counter++
	})
	return lazyInstance
}

var once sync.Once

func init() {
	hungryInstance = new(Instance)
	hungryInstance.Counter++
}

type Instance struct{ Counter int }
