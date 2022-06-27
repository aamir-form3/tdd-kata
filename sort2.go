package tdd_kata

func QSort(list []int) []int {
	result := make([]int, 0)
	if len(list) == 0 {
		return result
	}

	var less []int
	var greater []int
	middle := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] <= middle {
			less = append(less, list[i])
		} else {
			greater = append(greater, list[i])
		}
	}
	result = append(result, QSort(less)...)
	result = append(result, middle)
	result = append(result, QSort(greater)...)
	return result
}
