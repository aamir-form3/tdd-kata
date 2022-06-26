package tdd_kata

func Sort(nums []int) []int {

	if len(nums) > 1 {
		for limit := len(nums) - 1; limit > 0; limit-- {

			for first := 0; first < limit; first++ {
				second := first + 1
				if nums[first] > nums[second] {
					temp := nums[first]
					nums[first] = nums[second]
					nums[second] = temp
				}
			}
		}
	}
	return nums
}
