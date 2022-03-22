package main

import (
	"patterns/singleton"
	"patterns/builder"
	"patterns/objectpool"
	"time"
)

func main() {
	b := builder.NewCalc().Add(5).Del(2).Multi(4).Dev(3).Add(3).Build()
	b.Echo()

	st := singleton.NewInstance(5)
	st.Echo()

	st2 := singleton.NewInstance(6)
	st2.Echo()
	
	p := objectpool.NewPool(3)
loop:
	for {
		select {
		case r := <-p:
			r.Echo()
			p <- r
		default:
			break loop
		}
		time.Sleep(1 * time.Second)
	}
}
