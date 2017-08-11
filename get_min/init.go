package get_min

import (
	"data_structure/stack"
	"fmt"
)

type Min struct {
}

var stackData = stack.NewStack()
var stackMin = stack.NewStack()

func (min *Min) Pop() interface{} {
	num := stackData.Pop()
	if num == min.GetMin() {
		return stackMin.Pop()
	}
	return nil
}

func (min *Min) Push(value interface{}) {
	if stackMin.Empty() {
		stackMin.Push(value)
	} else if value < min.GetMin() {
		stackMin.Push(value)
	}
	stackData.Push(value)
}

func (min *Min) GetMin() interface{} {
	if stackMin.Empty() {
		fmt.Println("空栈")
	}
	return stackMin.Peak()
}
