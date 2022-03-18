package builder

import "fmt"

//产品可以拆分成多个部件,多个部件分别进行构造,最终行成的产品又具有相同的方法
//产品部件属性在创建完成之后,不可更改


// 产品接口定义
type Product interface {
	Echo()
}

// Builder接口定义
type Builder interface {
	Add(i int) Builder
	Del(i int) Builder
	Dev(i int) Builder
	Multi(i int) Builder

	Build() Product
}

// calc Builder实现
type Calc struct {
	value int
}

func (c *Calc) Add(i int) Builder {
	c.value += i
	return c
}
func (c *Calc) Del(i int) Builder {
	c.value -= i
	return c
}
func (c *Calc) Dev(i int) Builder {
	c.value = c.value / i
	return c
}
func (c *Calc) Multi(i int) Builder {
	c.value = c.value * i
	return c
}

func (c *Calc) Build() Product {
	var p MyProduct
	p.value = c.value
	return p
}

// 产品实现
type MyProduct struct {
	value int
}

func (p MyProduct) Echo() {
	fmt.Println(p.value)

}
func NewCalc() Builder {
	var c Calc
	return &c
}
