package main

import (
	"reflect"
	"testing"
)

// 参考:https://www.liwenzhou.com/posts/Go/unit-test/
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

// [2,4,9]
// [5,6,4,9]
// [7,0,4,0,1]
func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{Val: 9}

	except := &ListNode{Val: 7}
	except.Next = &ListNode{Val: 0}
	except.Next.Next = &ListNode{Val: 4}
	except.Next.Next.Next = &ListNode{Val: 0}
	except.Next.Next.Next.Next = &ListNode{Val: 1}

	got := addTwoNumbers(l1, l2)
	equal := true
	for got != nil && except != nil {
		if got.Val != except.Val {
			equal = false
			break
		}
		got = got.Next
		except = except.Next
	}
	if !equal {
		t.Errorf("addTwoNumbers expected be [7,0,4,0,1], but %v got", got)
	}
}

func TestSetZeroes(t *testing.T) {
	nums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	target := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	setZeroes(nums)
	if !reflect.DeepEqual(nums, target) {
		t.Errorf("TwoSum expected be %v, but %v got", nums, target)
	}
}

// go test -v
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		want  string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"len6":  {input: "1-23-45 6", want: "123-456"},
		"len7":  {input: "123 4-567", want: "123-45-67"},
		"len8":  {input: "123 4-5678", want: "123-456-78"},
		"len2":  {input: "12", want: "12"},
		"len13": {input: "--17-5 229 35-39475 ", want: "175-229-353-94-75"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := reformatNumber(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// go test -v
func TestCanTransform(t *testing.T) {
	type test struct { // 定义test结构体
		start string
		end   string
		want  bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"case1": {start: "XRXXXLXXXR", end: "XXRLXXXRXX", want: false},
		"case2": {start: "XXXXXLXXXLXXXX", end: "XXLXXXXXXXXLXX", want: false},
		"case3": {start: "RXXLRXRXL", end: "XRLXXRRLX", want: true},
		"case4": {start: "X", end: "L", want: false},
		"case5": {start: "RL", end: "LR", want: false},
		"case6": {start: "RLX", end: "XLR", want: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := canTransform(tc.start, tc.end)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
