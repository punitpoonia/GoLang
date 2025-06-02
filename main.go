package main

import (
	"fmt"
)

var stack []int

var anyDataTypeStack []interface{}

func pushAnyData(value interface{}) {
	anyDataTypeStack = append(anyDataTypeStack, value)
	fmt.Println("Pushed (value):", value)
}

func popFromAnyDataTypeStack() {
	if len(anyDataTypeStack) == 0 {
		fmt.Println("stack is empty. Cannot perform pop operation.")
		return
	}
	lastElement := anyDataTypeStack[len(anyDataTypeStack)-1]
	anyDataTypeStack = anyDataTypeStack[:len(anyDataTypeStack)-1]
	fmt.Println("Popped Element:", lastElement)
}

func push(value int) {
	stack = append(stack, value)
	fmt.Println("Pushed : ", value)
}

func pop() {

	if len(stack) == 0 {
		fmt.Println("Stack is empty. Cannot perform pop operation.")
		return
	}
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
	pop()
	pop()
	pop()
	fmt.Println("After Pop Operation ", stack)

	// AnyDataType Stack Operations
	pushAnyData(42)
	pushAnyData("GoLang")
	pushAnyData(3.14)
	pushAnyData(true)
	pushAnyData([]int{1, 2, 3})
	pushAnyData(map[string]int{"a": 1})

	fmt.Println("Any Data Type Stack:", anyDataTypeStack)

	popFromAnyDataTypeStack()
	popFromAnyDataTypeStack()

	fmt.Println("After Pop Operation in AnyDataTypeStack : ",anyDataTypeStack)

}
