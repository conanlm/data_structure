package main

import (
	"data_structure/get_min"
	"fmt"
)

func main() {
	data := get_min.New()
	data.Push(10)
	data.Push(1)
	data.Push(0)
	fmt.Println(data.GetMin())
}
