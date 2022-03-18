package main

import "patterns/builder"

func main() {
	b := builder.NewCalc().Add(5).Del(2).Multi(4).Dev(3).Add(3).Build()
	b.Echo()

	
}
