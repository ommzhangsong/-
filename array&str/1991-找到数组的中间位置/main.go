package main

// 如下：暴力解法
func sum(arr []int) int {
	var result int
	for _, v := range arr {
		result += v
	}
	return result
}
func findMiddleIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if sum(nums[:i]) == sum(nums[i+1:]) {
			return i
		}
	}
	return -1
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	findMiddleIndex(nums)
}
