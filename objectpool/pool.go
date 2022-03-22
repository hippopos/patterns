package objectpool

import "fmt"

type Pool struct {
	value int
}

func (p Pool) Echo() {
	fmt.Println("echo lalala.", p.value)
}
func NewPool(count int) chan Pool {
	pc := make(chan Pool, count)
	for i := 0; i < count; i++ {
		pc <- Pool{value: i}
	}
	return pc
}
