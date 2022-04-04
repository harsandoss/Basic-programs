package main

import (
	"fmt"
)

func add(x, y int) int {
	var a = x + y
	return a
}
func sub(x, y int) int {
	var s = x - y
	return s
}
func mul(x, y int) int {
	var m = x * y
	return m
}
func div(x, y int) int {
	var d = x / y
	return d
}
func arithematic(x, y int) (add int, s int, m int, d int) {
	add = x + y

	s = x - y

	m = x * y

	d = x / y
	return

}
func main() {
	add := add(3, 4)
	fmt.Println(add)
	sub := sub(12, 4)
	fmt.Println(sub)
	mul := mul(3, 4)
	fmt.Println(mul)
	div := div(12, 4)
	fmt.Println(div)
	var a, s, m, d = arithematic(10, 5)
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(d)
}
