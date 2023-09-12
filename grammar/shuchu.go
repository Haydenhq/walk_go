package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	// return fmt.Sprintf("print: %v", c)
	return "we"
}

func main() {
	c := &ConfigOne{}
	c.String()

	s := getP()
	fmt.Println(&s) // cannot take the address of getP()
	// fmt.Println(&getP()) // cannot take the address of getP()
}

type person struct {
	name string
}

func getP() person {
	return person{name: "yif"}
}
