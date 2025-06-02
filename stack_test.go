package main

import (
	"reflect"
	"testing"
)

func TestStackOfIntPush(t *testing.T) {
	stack := StackOfInt{}
	stack.push(1)
	stack.push(2)

	expected := []int{1, 2}
	if !reflect.DeepEqual(stack.structInput, expected) {
		t.Errorf("Expected %v, got %v", expected, stack.structInput)
	}
}

func TestStackOfIntPop(t *testing.T) {
	stack := StackOfInt{}
	stack.push(10)
	stack.push(20)
	stack.pop()

	expected := []int{10}
	if !reflect.DeepEqual(stack.structInput, expected) {
		t.Errorf("Expected %v after pop, got %v", expected, stack.structInput)
	}
}

func TestStackOfIntPopEmpty(t *testing.T) {
	stack := StackOfInt{}
	stack.pop() // Should not crash

	if len(stack.structInput) != 0 {
		t.Errorf("Expected empty stack, got %v", stack.structInput)
	}
}

func TestAnyDataTypeStackPush(t *testing.T) {
	stack := AnyDataTypeStack{}
	stack.pushAnyData(42)
	stack.pushAnyData("hello")
	stack.pushAnyData(3.14)

	expected := []interface{}{42, "hello", 3.14}
	if !reflect.DeepEqual(stack.anyDataType, expected) {
		t.Errorf("Expected %v, got %v", expected, stack.anyDataType)
	}
}

func TestAnyDataTypeStackPop(t *testing.T) {
	stack := AnyDataTypeStack{}
	stack.pushAnyData("first")
	stack.pushAnyData("second")
	stack.popFromAnyDataTypeStack()

	expected := []interface{}{"first"}
	if !reflect.DeepEqual(stack.anyDataType, expected) {
		t.Errorf("Expected %v after pop, got %v", expected, stack.anyDataType)
	}
}

func TestAnyDataTypeStackPopEmpty(t *testing.T) {
	stack := AnyDataTypeStack{}
	stack.popFromAnyDataTypeStack()

	if len(stack.anyDataType) != 0 {
		t.Errorf("Expected empty stack, got %v", stack.anyDataType)
	}
}
