package main

import "fmt"

func Foo(x any) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {

	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	_ = y == y

	var x1 []int = nil
	Foo(x1)
}
