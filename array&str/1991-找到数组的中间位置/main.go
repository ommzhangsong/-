package main

// 如下：暴力解法,时空复杂度，n平方，和常数
//
//	func sum(arr []int) int {
//		var result int
//		for _, v := range arr {
//			result += v
//		}
//		return result
//	}
//
//	func findMiddleIndex(nums []int) int {
//		for i := 0; i < len(nums); i++ {
//			if sum(nums[:i]) == sum(nums[i+1:]) {
//				return i
//			}
//		}
//		return -1
//	}
func findMiddleIndex(nums []int) int {
	var sum, lsum int
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		// 这里注意lsum和i的关系，我们先判断，代表i=0时，开始此时lsum=0，符合题意
		// 如果先算lsum那么加了第一个值，那么按道理说是按i=1来的，但返回任然是i不对，
		// 必须先从i=0开始判断，所以return i也是不能变的
		if lsum*2 == sum-nums[i] {
			return i
		}
		lsum += nums[i]
	}
	return -1
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	findMiddleIndex(nums)
}
