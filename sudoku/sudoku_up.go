package main

import (
	"container/list"
	"fmt"
	"os"
)

type Recoder struct {
	point      [2]int
	pointIndex int
	screen     map[int][]int
}

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
			sudo.Screen(point[0]*9+key, [2]int{point[0], key}, val)
		}
	}

	//列排除
	for i := 0; i < 9; i++ {
		if sudo.value[i][point[1]] == 0 {
			sudo.Screen(i*9+point[1], [2]int{i, point[1]}, val)
		}
	}

	//九宫格排除
	x := point[0] / 3 * 3
	y := point[1] / 3 * 3
	for key, _ := range sudo.value[x : x+3] {
		for i := y; i < y+3; i++ {
			if sudo.value[key][i] == 0 {
				sudo.Screen(key*3+i, [2]int{key, i}, val)
			}
		}
	}
}

func (sudo *Sudo) Screen(key int, point [2]int, block int) {
	list := sudo.screen[key]
	for k, col := range list {
		if block == col {
			sudo.screen[key] = append(list[:k], list[k+1:]...)
		}
	}
	if len(sudo.screen[key]) == 1 {
		sudo.new_points.PushFront(point)
		sudo.value[point[0]][point[1]] = sudo.screen[key][0]
	}
}

func (sudo *Sudo) CheckOnePossbile() bool {
	for r := range [9]int{0: 9} {
		for k, c := range sudo.value[r] {
			if sudo.value[r][k] == 0 {

			}
		}
	}
	return true
}

func (sudo *Sudo) CheckSameNum() {
	for _, val := range sudo.base_points {
		for key, _ := range sudo.value[val[0] : val[0]+3] {
			for i := val[1]; i < val[1]+3; i++ {
				if sudo.value[key][i] == 0 {

				}
			}
		}
	}
}

//得到确定的数字
func (sudo *Sudo) GetCount() int {
	sum := 0
	for i := 0; i < 81; i++ {
		if sudo.value[i/9][i%9] != 0 {
			sum++
		}
	}
	return sum
}

//评分，找到最佳的猜测坐标
func (sudo *Sudo) GetBestPoint() [2]int {
	bestScore := 0
	bestPoint := [2]int{0, 0}
	for row, _ := range sudo.value {
		for col, _ := range sudo.value[row] {
			pointScore := sudo.getPointScore([2]int{row, col})
			if bestScore < pointScore {
				bestScore = pointScore
				bestPoint = [2]int{row, col}
			}
		}
	}
	return bestPoint
}

//计算某坐标的评分
func (sudo *Sudo) getPointScore(point [2]int) int {
	if sudo.value[point[0]][point[1]] == 0 {
		score := 10 - len(sudo.screen[point[0]*9+point[1]])
		for _, val := range sudo.value[point[0]] {
			if val > 0 {
				score++
			}
		}
		for i := 0; i < 9; i++ {
			if sudo.value[i][point[1]] > 0 {
				score++
			}
		}
		return score
	}
	return 0
}

//检查数字有无错误
func (sudo *Sudo) CheckValue() bool {
	for row, _ := range sudo.value {
		nums := make([]int, 0)
		lists := make([][]int, 0)
		for col, val := range sudo.value[row] {
			if val == 0 {
				lists = append(lists, sudo.screen[row*9+col])
			} else {
				nums = append(nums, val)
			}
		}
		if isRetBol(nums, lists) == false {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		nums := make([]int, 0)
		lists := make([][]int, 0)
		for j := 0; j < 9; j++ {
			if sudo.value[j][i] == 0 {
				lists = append(lists, sudo.screen[j*9+i])
			} else {
				nums = append(nums, sudo.value[j][i])
			}
		}
		if isRetBol(nums, lists) == false {
			return false
		}
	}

	for _, val := range sudo.base_points {
		for key, _ := range sudo.value[val[0] : val[0]+3] {
			nums := make([]int, 0)
			lists := make([][]int, 0)
			for i := val[1]; i < val[1]+3; i++ {
				if sudo.value[key][i] == 0 {
					lists = append(lists, sudo.screen[key*9+i])
				} else {
					nums = append(nums, sudo.value[key][i])
				}
			}
			if isRetBol(nums, lists) == false {
				return false
			}
		}
	}
	return true
}

//猜测记录
func (sudo *Sudo) RecodeGuess(point [2]int, index int) {
	var recoder Recoder
	recoder.point = point
	recoder.pointIndex = index
	recoder.screen = sudo.screen
	sudo.recoder.PushFront(recoder)
	sudo.guess_times++

	//新一轮的排除处理
	item := sudo.screen[point[0]*9+point[1]]
	sudo.value[point[0]][point[1]] = item[index]
	sudo.new_points.PushFront(point)
	sudo.SolveSudo()
}

//回溯，需要先进先出
func (sudo *Sudo) Reback() {
	var index int
	var point [2]int
	var recoder Recoder
	for {
		if sudo.recoder.Len() == 0 {
			fmt.Println("sudo is wrong")
			os.Exit(0)
		} else {
			recoder = sudo.recoder.Back().Value.(Recoder)
			point = recoder.point
			index = recoder.pointIndex + 1
			item := recoder.screen[point[0]*9+point[1]]
			if index < len(item) {
				break
			}
		}
	}
	sudo.screen = recoder.screen
	sudo.RecodeGuess(point, index)
}

func isRetBol(nums []int, lists [][]int) bool {
	if len(removeDuplicates(nums)) != len(nums) {
		return false
	}
	for _, val := range lists {
		if len(val) == 0 {
			return false
		}
	}
	return true
}

func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func main() {
	data := New()
	// data.Calc()
	// fmt.Println(append(data.screen[0][:3], data.screen[0][3+2:]...))
	// fmt.Println(data.new_points.Front().Value)
	// data.CutNum([2]int{0, 7})
	// data.CheckSameNum()
	data.CheckValue()
	fmt.Println("")
}
