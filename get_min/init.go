package get_min

import (
	"data_structure/stack"
	"fmt"
)

type Min struct {
	stackData, stackMin *stack.Stack
}

func New() *Min {
	stackData := stack.NewStack()
	stackMin := stack.NewStack()
	return &Min{stackData, stackMin}
}

func (min *Min) Pop() interface{} {
	if min.stackData.Empty() {
		fmt.Println("空栈")
	}
	num := min.stackData.Pop()
	if num == min.GetMin() {
		return min.stackMin.Pop()
	}
	return nil
}

func (min *Min) Push(value interface{}) {
	if min.stackMin.Empty() {
		min.stackMin.Push(value)
	}
	min.stackData.Push(value)
	if value.(int) <= min.GetMin().(int) {
		min.stackMin.Push(value)
	}
}

func (min *Min) GetMin() interface{} {
	if min.stackMin.Empty() {
		fmt.Println("空栈")
	}
	return min.stackMin.Peak()
}
