package main

import (
	"container/list"
	"fmt"
	"github.com/mohae/deepcopy"
	"os"
)

type Recoder struct {
	point      [2]int
	pointIndex int
	screen     map[int][]int
	sudokuList [9][9]interface{}
}

type Sudo struct {
	guess_times int
	value       [9][9]int
	new_points  *list.List
	recoder     *list.List
	base_points [9][2]int
	screen      map[int][]int
	sudokuList  [9][9]interface{}
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
	var sudokuList [9][9]interface{}
	new_points := list.New()
	recoder := list.New()
	for i := 0; i < 81; i++ {
		if sudoArr[i/9][i%9] != 0 {
			sudokuList[i/9][i%9] = sudoArr[i/9][i%9]
			new_points.PushBack([2]int{i / 9, i % 9})
		} else {
			sudokuList[i/9][i%9] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		}
	}
	base_points := [9][2]int{
		{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6},
	}
	return &Sudo{base_points: base_points, guess_times: 0,
		new_points: new_points, recoder: recoder, sudokuList: sudokuList}
}

func (sudo *Sudo) Calc() {
	sudo.SolveSudo()
	for {
		if sudo.CheckValue() {
			if sudo.GetCount() == 81 {
				break
			}
			point := sudo.GetBestPoint()
			sudo.RecodeGuess(point, 0)
		} else {
			sudo.Reback()
		}
	}
}

func (sudo *Sudo) SolveSudo() {
	isRunSame := true
	isRunOne := true
	for isRunSame {
		for isRunOne {
			for sudo.new_points.Len() > 0 {
				e := sudo.new_points.Front()
				sudo.new_points.Remove(e)
				if sudo.new_points.Len()==9{
					fmt.Println(sudo.sudokuList[0])
				}
				sudo.CutNum(e.Value.([2]int))
				if sudo.new_points.Len()==9{
					fmt.Println(sudo.sudokuList[0])
				}
			}
			os.Exit(1)
			isRunOne = sudo.CheckOnePossbile()
		}
		isRunSame = sudo.CheckSameNum()
		isRunOne = true
	}
}

func (sudo *Sudo) CutNum(point [2]int) {
	val,err:= sudo.sudokuList[point[0]][point[1]].(int)
	if !err{
		return
	}
	if sudo.new_points.Len()==9{
		fmt.Println("")
	}
	//行排除
	for col:= range [9]int{0:9} {
		item,ok:= sudo.sudokuList[point[0]][col].([]int)
		if !ok {
			continue
		}
		key:=In(val, item)
		if key==-1{
			continue
		}
		if sudo.new_points.Len()==9{
			fmt.Println(point[0],col)
			fmt.Println(sudo.sudokuList[0])
		}
		temp:=CopySlice(item, key)
		sudo.sudokuList[point[0]][col]=temp
		if len(temp)==1{
			sudo.new_points.PushFront([2]int{point[0],col})
			sudo.sudokuList[point[0]][col]=temp[0]
		}
	}


	//列排除
	for row:=range [9]int{0:9}{
		item,ok:=sudo.sudokuList[row][point[1]].([]int)
		if !ok {
			continue
		}
		key:=In(val, item)
		if key==-1{
			continue
		}
		temp:=CopySlice(item, key)
		sudo.sudokuList[row][point[1]]=temp
		if len(temp)==1{
			sudo.new_points.PushFront([2]int{row,point[1]})
			sudo.sudokuList[row][point[1]]=temp[0]
		}
	}

	//九宫格排除
	x := point[0] / 3 * 3
	y := point[1] / 3 * 3
	for row, _ := range sudo.sudokuList[x : x+3] {
		for col := y; col < y+3; col++ {
			if _, ok := sudo.sudokuList[row][col].([]int); !ok {
				continue
			}
			key:=In(val, sudo.sudokuList[row][col].([]int))
			if key==-1{
				continue
			}
			temp:=CopySlice(sudo.sudokuList[row][col].([]int), key)
			sudo.sudokuList[row][col]=temp
			if len(temp)==1{
				sudo.new_points.PushFront([2]int{row,col})
				sudo.sudokuList[row][col]=temp[0]
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
		for k, _ := range sudo.sudokuList[r] {
			if _, ok := sudo.sudokuList[r][k].([]int); !ok {
				continue
			}
			for _, val := range sudo.sudokuList[r][k].([]int) {
				sum := sudo.ergodic(sudo.sudokuList[r][k].([]int), r, val)
				if sum == 1 {
					sudo.sudokuList[r][k] = val
					sudo.new_points.PushFront([2]int{r, k})
					return true
				}
			}
		}
	}

	for c := range [9]int{0: 9} {
		for r, _ := range sudo.sudokuList[:][c] {
			if _, ok := sudo.sudokuList[r][c].([]int); !ok {
				continue
			}
			for _, val := range sudo.sudokuList[r][c].([]int) {
				sum := sudo.ergodic(sudo.sudokuList[r][c].([]int), r, val)
				if sum == 1 {
					sudo.sudokuList[r][c] = val
					sudo.new_points.PushFront([2]int{r, c})
					return true
				}
			}
		}
	}

	for _, val := range sudo.base_points {
		for key, _ := range sudo.sudokuList[val[0] : val[0]+3] {
			for i := val[1]; i < val[1]+3; i++ {
				if _, ok := sudo.sudokuList[key][i].([]int); !ok {
					continue
				}

				for _, v := range sudo.sudokuList[key][i].([]int) {
					sum := sudo.ergodic(sudo.sudokuList[key][i].([]int), key, v)
					if sum == 1 {
						sudo.value[key][i] = v
						sudo.new_points.PushFront([2]int{key, i})
						return true
					}
				}
			}
		}
	}
	return false
}

func (sudo *Sudo) ergodic(list []int, row int, search int) int {
	sum := 0
	for key, _ := range list {
		if _, ok := sudo.sudokuList[row][key].([]int); !ok {
			continue
		}
		for _, val := range sudo.sudokuList[row][key].([]int) {
			if search == val {
				sum++
			}
		}
	}
	return sum
}

func (sudo *Sudo) CheckSameNum() bool {
	for _, val := range sudo.base_points {
		for i := 1; i < 10; i++ {
			result := make([]int, 0)
			for key, _ := range sudo.sudokuList[val[0] : val[0]+3] {
				for j := val[1]; j < val[1]+3; j++ {
					if _, ok := sudo.sudokuList[key][j].([]int); !ok {
						continue
					}
					if blockKey := In(i, sudo.sudokuList[key][j].([]int)); blockKey > -1 {
						result = append(result, blockKey)
					}
				}

			}
			if rCount := len(result); rCount == 2 || rCount == 3 {
				rows := Map(func(x int) int { return x / 3 }, result)
				cols := Map(func(x int) int { return x % 3 }, result)
				if len(removeDuplicates(rows)) == 1 {
					row := val[0] + rows[0]
					result = Map(func(x int) int { return val[0] + x%3 }, result)

					for col := range [9]int{0: 9} {
						if In(col, result) > -1 {
							continue
						}
						if _, ok := sudo.sudokuList[row][col].([]int); !ok {
							continue
						}
						if sudo.replace(row, col, i) {
							return true
						}
					}
				} else if len(removeDuplicates(cols)) == 1 {
					result = Map(func(x int) int { return val[0] + x/3 }, result)
					col := val[1] + cols[0]

					for row := range [9]int{0: 9} {
						if In(row, result) > -1 {
							continue
						}
						if _, ok := sudo.sudokuList[row][col].([]int); !ok {
							continue
						}
						if sudo.replace(row, col, i) {
							return true
						}
					}
				}

			}
		}
	}
	return false
}

func (sudo *Sudo) replace(row int, col int, search int) bool {
	if _, ok := sudo.sudokuList[row][col].([]int); !ok {
		return false
	}
	key := In(search, sudo.sudokuList[row][col].([]int))
	if key == -1 {
		return false
	}
	temp := deepcopy.Copy(sudo.sudokuList[row][col].([]int)).([]int)
	temp = append(temp[:key], temp[key+1:]...)
	sudo.sudokuList[row][col] = temp
	if len(temp) == 1 {
		// sudo.new_points.PushFront([2]int{row, col})
		sudo.sudokuList[row][col] = temp[0]
		return true
	}
	return false
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

//得到确定的数字
func (sudo *Sudo) GetCount() int {
	sum := 0
	for i := 0; i < 81; i++ {
		if _, ok := sudo.sudokuList[i/9][i%9].(int); ok {
			sum++
		}
	}
	return sum
}

//评分，找到最佳的猜测坐标
func (sudo *Sudo) GetBestPoint() [2]int {
	bestScore := 0
	bestPoint := [2]int{0, 0}
	for row := range [9]int{0: 9} {
		for col, _ := range sudo.sudokuList[row] {
			pointScore := sudo.getPointScore(row, col)
			if bestScore < pointScore {
				bestScore = pointScore
				bestPoint = [2]int{row, col}
			}
		}
	}
	return bestPoint
}

//计算某坐标的评分
func (sudo *Sudo) getPointScore(row int, col int) int {
	if _, ok := sudo.sudokuList[row][col].([]int); !ok {
		return 0
	}
	score := 10 - len(sudo.sudokuList[row][col].([]int))
	for _, val := range sudo.sudokuList[row] {
		if _, ok := val.(int); ok {
			score++
		}
	}
	for i := 0; i < 9; i++ {
		if _, ok := sudo.sudokuList[i][col].(int); ok {
			score++
		}
	}
	return score
}

//检查数字有无错误
func (sudo *Sudo) CheckValue() bool {
	for row, _ := range sudo.sudokuList {
		nums := make([]int, 0)
		lists := make([][]int, 0)
		for _, val := range sudo.sudokuList[row] {
			if list, ok := val.([]int); ok {
				lists = append(lists, list)
			} else {
				nums = append(nums, val.(int))
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
			if list, ok := sudo.sudokuList[j][i].([]int); ok {
				lists = append(lists, list)
			} else {
				nums = append(nums, sudo.sudokuList[j][i].(int))
			}
		}
		if isRetBol(nums, lists) == false {
			return false
		}
	}

	for _, val := range sudo.base_points {
		for key, _ := range sudo.sudokuList[val[0] : val[0]+3] {
			nums := make([]int, 0)
			lists := make([][]int, 0)
			for i := val[1]; i < val[1]+3; i++ {
				if list, ok := sudo.sudokuList[key][i].([]int); ok {
					lists = append(lists, list)
				} else {
					nums = append(nums, sudo.sudokuList[key][i].(int))
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
	recoder.sudokuList = sudo.sudokuList
	sudo.recoder.PushFront(recoder)
	sudo.guess_times++

	//新一轮的排除处理
	item := sudo.sudokuList[point[0]][point[1]].([]int)
	sudo.sudokuList[point[0]][point[1]] = item[index]
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
			e := sudo.recoder.Back()
			sudo.recoder.Remove(e)
			recoder = e.Value.(Recoder)
			point = recoder.point
			index = recoder.pointIndex + 1
			item := recoder.sudokuList[point[0]][point[1]].([]int)
			if index < len(item) {
				break
			}
		}
	}
	sudo.sudokuList = recoder.sudokuList
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
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] != true {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func Map(f func(int) int, v []int) (r []int) {
	r = make([]int, len(v))
	for i, value := range v {
		r[i] = f(value)
	}
	return
}

func In(search int, value []int) int {
	for k, val := range value {
		if search == val {
			return k
		}
	}
	return -1
}

func CopySlice(arr []int, key int)[]int{
	list:=make([]int, len(arr)-1)
	i:=0
	for k,val:=range arr{
		if k==key{
			continue
		}
		list[i]=val
		i++
	}
	return list
}

func main() {
	data := New()
	data.Calc()
	fmt.Println(data.sudokuList)
	fmt.Println("")
}
