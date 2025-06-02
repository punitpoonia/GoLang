package main

import (
	"fmt"
)

type StackOfInt struct {
	structInput []int
}

type AnyDataTypeStack struct {
	anyDataType []interface{}
}

func (d *AnyDataTypeStack) pushAnyData(value interface{}) {
	d.anyDataType = append(d.anyDataType, value)
	fmt.Println("Pushed (value):", value)
}

func (d *AnyDataTypeStack) popFromAnyDataTypeStack() {
	if len(d.anyDataType) == 0 {
		fmt.Println("stack is empty. Cannot perform pop operation.")
		return
	}
	lastElement := d.anyDataType[len(d.anyDataType)-1]
	d.anyDataType = d.anyDataType[:len(d.anyDataType)-1]
	fmt.Println("Popped Element:", lastElement)
}

func (data *StackOfInt) push(value int) {
	data.structInput = append(data.structInput, value)
	fmt.Println("Pushed : ", value)
}

func (data *StackOfInt) pop() {

	if len(data.structInput) == 0 {
		fmt.Println("Stack is empty. Cannot perform pop operation.")
		return
	}
	lastElemet := data.structInput[len(data.structInput)-1]
	data.structInput = data.structInput[:len(data.structInput)-1]

	fmt.Println("Poped Element : ", lastElemet)
}

func main() {
	intStack := StackOfInt{}
	intStack.push(1)
	intStack.push(2)

	fmt.Println(intStack.structInput)
	intStack.pop()
	intStack.pop()
	intStack.pop()

	fmt.Println("After Pop Operation ", intStack.structInput)

	// AnyDataType Stack Operations
	anyStack := AnyDataTypeStack{}
	anyStack.pushAnyData(42)
	anyStack.pushAnyData("GoLang")
	anyStack.pushAnyData(3.14)
	anyStack.pushAnyData([]int{1, 2, 3})
	anyStack.pushAnyData(map[string]int{"a": 1})

	fmt.Println("Any Data Type Stack:", anyStack.anyDataType)

	anyStack.popFromAnyDataTypeStack()
	anyStack.popFromAnyDataTypeStack()

	fmt.Println("After Pop Operation in AnyDataTypeStack : ", anyStack.anyDataType)

}
