package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var i int = len(nums)
	var x int = 0
	var y int = 0
	var sum int = 0
	for x = 0; x <= i-1; x++ {
		y = x + 1
		if i != 1 {
			for ; y <= i-1; y++ {
				sum = nums[x] + nums[y]
				if sum == target {
					i = 1
					break
				}

			}
			if i == 1 {
				break
			}
		}

	}
	return []int{x, y}
}

func main() {
	var nums = []int{3, 2, 4}
	var target int = 6
	fmt.Println(twoSum(nums, target))
}
