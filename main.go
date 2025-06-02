package main

import "fmt"

var stack []int

func push(value int) {
	stack = append(stack, value)
	fmt.Println("Pushed : ", value)
}

func pop() {
	lastElemet := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	fmt.Println("Poped Element : ", lastElemet)
}

func main() {
	push(1)
	push(2)
	push(3)
	fmt.Println(stack)
	pop()
	fmt.Println("After Pop Operation ", stack)
}
