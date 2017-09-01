package sort

import (
	"testing"
)

var list = []int{20, 10, 5, 56, 21, 33}

func TestBubbleSort1(t *testing.T) {
	t.Log(BubbleSort1(list))
}

func TestBubbleSort2(t *testing.T) {
	t.Log(BubbleSort2(list))
}

func TestBubbleSort3(t *testing.T) {
	t.Log(BubbleSort3(list))
}
