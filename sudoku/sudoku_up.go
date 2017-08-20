package main

import (
	"container/list"
	"fmt"
)

type Sudo struct {
	guess_times int
	value       [9][9]int
	new_points  *list.List
	recoder     *list.List
	base_points [9][2]int
	screen      map[int][]int
}

func New() *Sudo {
	sudoArr := [9][9]int{
		{0, 0, 0, 0, 0, 2, 0, 5, 0},
		{0, 7, 8, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 0, 0, 4, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 3, 0, 7, 0, 8},
		{2, 0, 0, 0, 0, 0, 0, 4, 0},
		{0, 0, 0, 0, 0, 5, 0, 9, 0},
		{0, 1, 0, 0, 7, 0, 0, 0, 0},
	}
	new_points := list.New()
	recoder := list.New()
	screen := make(map[int][]int, 0)
	for i := 0; i < 81; i++ {
		if sudoArr[i/9][i%9] != 0 {
			new_points.PushBack([2]int{i / 9, i % 9})
		} else {
			screen[i] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		}
	}
	base_points := [9][2]int{
		{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6},
	}
	return &Sudo{value: sudoArr, base_points: base_points, guess_times: 0,
		new_points: new_points, recoder: recoder, screen: screen}
}

func (sudo *Sudo) Calc() {
	sudo.SolveSudo()
}

func (sudo *Sudo) SolveSudo() {
	isRunSame := true
	isRunOne := true
	var point [2]int
	for isRunSame {
		for isRunOne {
			for sudo.new_points.Len() > 0 {
				point = sudo.new_points.Front().Value.([2]int)
				sudo.CutNum(point)
			}
		}

		isRunOne = true
	}
}

func (sudo *Sudo) CutNum(point [2]int) {
	val := sudo.value[point[0]][point[1]]
	//行排除
	for key, row := range sudo.value[point[0]] {
		if row == 0 {
			screenKey := point[0]*9 + key
			list := sudo.screen[screenKey]
			for k, col := range list {
				if val == col {
					sudo.screen[screenKey] = append(list[:k], list[k+1:]...)
				}
			}
			if len(sudo.screen[screenKey]) == 1 {
				sudo.new_points.PushFront(point)
				sudo.value[point[0]][point[1]] = sudo.screen[screenKey][0]
			}
		}
	}
}

func (sudo *Sudo) CheckSameNum() {
	// for _, val := range sudo.base_points {

	// }
}

func main() {
	data := New()
	// data.Calc()
	// fmt.Println(append(data.screen[0][:3], data.screen[0][3+2:]...))
	fmt.Println(data.new_points.Front().Value)
	data.CutNum([2]int{0, 7})
}
