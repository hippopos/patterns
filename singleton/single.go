package singleton

import (
	"fmt"
	"sync"
)

type instance struct {
	value int
}

func (inst instance) Echo() {
	fmt.Println(inst.value)
}

var inst instance
var once sync.Once

func NewInstance(v int) instance {
	once.Do(
		func() {
			inst = instance{value: v}
		})
	return inst
}
