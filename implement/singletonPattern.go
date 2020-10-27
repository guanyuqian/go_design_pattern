package implement

import "sync"

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
