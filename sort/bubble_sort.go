package sort

func BubbleSort1(list []int) ([]int, int) {
	count := 0
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			count++
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list, count
}

func BubbleSort2(list []int) ([]int, int) {
	count := 0
	for i := 1; i < len(list); i++ {
		for j := 0; j < len(list)-i-1; j++ {
			count++
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list, count
}

func BubbleSort3(list []int) ([]int, int) {
	count := 0
	flag := true
	for i := 1; i < len(list) && flag; i++ {
		flag = false
		for j := 0; j < len(list)-i-1; j++ {
			count++
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
			}
		}
	}
	return list, count
}
