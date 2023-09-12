package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

// 指针类型的实现了  值的类型没有实现
func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// 值类型没有实现Speak 方法
	// var peo People = Student{}
	var peo1 People = &Student{}
	think := "speak"
	fmt.Println(peo1.Speak(think))
}
