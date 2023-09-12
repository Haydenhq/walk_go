package main

import (
	"fmt"
	"time"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		fmt.Println(i, v)
		fmt.Println(i, &v)
	}

	for s := range m {
		fmt.Println(s)
		fmt.Println(&s)
	}

	time.Sleep(time.Second * 3)
}
