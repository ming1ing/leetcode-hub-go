package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	if ans := twoSum(nums, target); ans[0] != 1 && ans[1] != 0 {
		t.Errorf("TwoSum expected be [1,0], but %v got", ans)
	}
	nums = []int{3, 2, 4}
	target = 6
	if ans := twoSum(nums, target); ans[0] != 2 && ans[1] != 1 {
		t.Errorf("TwoSum expected be [2,1], but %v got", ans)
	}
	nums = []int{3, 3}
	target = 6
	if ans := twoSum(nums, target); ans[0] != 1 && ans[1] != 0 {
		t.Errorf("TwoSum expected be [1,0], but %v got", ans)
	}
}
