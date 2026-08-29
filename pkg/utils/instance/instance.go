package instance

import (
	"sync"
)

type AppContext interface {
	Close() error
}

type Instance struct {
	instance  []AppContext
	closeOnce sync.Once
}

func New() *Instance {
	return &Instance{}
}

func (i *Instance) Add(instance AppContext) {
	i.instance = append(i.instance, instance)
}

func (i *Instance) Close() {
	i.closeOnce.Do(func() {
		for _, v := range i.instance {
			v.Close()
		}
	})
}
