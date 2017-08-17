package main

import (
    "fmt"
)


var Sudoku = [9][9]int{
    {0, 0, 0, 0, 0, 0, 8, 0, 0},
    {0, 8, 2, 4, 0, 0, 0, 0, 0},
    {1, 9, 0, 0, 6, 3, 0, 0, 0},
    {0, 5, 0, 0, 8, 0, 7, 0, 0},
    {6, 7, 8, 2, 0, 9, 1, 4, 3},
    {0, 0, 3, 0, 4, 0, 0, 8, 0},
    {0, 0, 0, 6, 2, 0, 0, 9, 4},
    {0, 0, 0, 0, 0, 5, 6, 1, 0},
    {0, 0, 0, 6, 0, 0, 0, 0, 0}}

func main() {
	fmt.Println("未填充数独：")
	Output(Sudoku)
}

func Output(sudoku [9][9]int) {
	for i := 0; i < 81; i++ {
	    fmt.Printf("%2d ", sudoku[i/9][i%9])
	    if i!=0 && (i+1)%9==0{
		    fmt.Println("")
	    }
	}
}
