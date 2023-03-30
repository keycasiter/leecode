package exchange

import (
	"fmt"
	"testing"
)

//变量交换
func TestExchange1(t *testing.T) {
	a := 1
	b := 2

	fmt.Printf("a:%d ,b:%d\n", a, b)

	tmp := a
	a = b
	b = tmp

	fmt.Printf("a:%d ,b:%d\n", a, b)
}

//加减法交换
func TestExchange2(t *testing.T) {
	a := 1
	b := 2

	fmt.Printf("a:%d ,b:%d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a:%d ,b:%d\n", a, b)
}

//异或交换
func TestExchange3(t *testing.T) {
	a := 1
	b := 2

	fmt.Printf("a:%d ,b:%d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a:%d ,b:%d\n", a, b)
}
