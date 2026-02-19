package main

import "fmt"

// 排序好的数组，去找索引位置，根据值
func searchInsert(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(searchInsert(nums, 5))
}
