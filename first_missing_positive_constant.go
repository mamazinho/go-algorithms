package main

import "fmt"

func FirstMissingPositive() {
	resp := Algorithm([]int{3, 4, -1, 1})
	fmt.Println(resp)
}

func Algorithm(nums []int) int {
	for i := 0; i < len(nums); {
		cur := nums[i]
		fmt.Println(nums, cur)
		if cur > 0 && cur <= len(nums) && nums[cur-1] != cur {
			nums[i], nums[cur-1] = nums[cur-1], nums[i]
			continue
		}
		i++
	}

	fmt.Println(nums)

	for i := 1; i <= len(nums)+1; i++ {
		if nums[i-1] != i {
			return i
		}
	}
	return len(nums) + 1
}
