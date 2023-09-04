package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	h := hello
	fmt.Printf("%T = %v", h)

	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
	fmt.Println(t[1])
	fmt.Println(len(t))
	fmt.Println(cap(t))

	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
