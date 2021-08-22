package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, 0)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if id, ok := numsMap[target-nums[i]]; ok {
			ans = append(ans, i)
			ans = append(ans, id)
			return ans
		}
		numsMap[nums[i]] = i
	}
	return ans
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))
	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
}
