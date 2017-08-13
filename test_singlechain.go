package main

import (
	"data_structure/singlechain"
	"fmt"
)

func main() {
	var h singlechain.Node
	for i := 1; i <= 10; i++ {
		var d singlechain.Node
		d.Data = i
		singlechain.Insert(&h, &d, i)
		fmt.Println(singlechain.GetLoc(&h, i))
	}
	fmt.Println(singlechain.GetLength(&h))
	fmt.Println(singlechain.GetFirst(&h))
	fmt.Println(singlechain.GetLast(&h))
	fmt.Println(singlechain.GetLoc(&h, 6))
}
