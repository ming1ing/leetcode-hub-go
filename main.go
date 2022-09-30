package main

import (
	"fmt"
	"strings"
)

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
func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	nn := make([]bool, n)
	mm := make([]bool, m)
	for i, v := range matrix {
		for j, v2 := range v {
			if v2 == 0 {
				nn[i] = true
				mm[j] = true
			}
		}
	}
	for k, v := range nn {
		if v == true {
			for i := 0; i < len(matrix[k]); i++ {
				matrix[k][i] = 0
			}
		}
	}
	for k, v := range mm {
		if v == true {
			for j := 0; j < len(matrix); j++ {
				matrix[j][k] = 0
			}
		}
	}
}
func reformatNumber(number string) string {
	ss := strings.Split(number, " ")
	ns := strings.Join(ss, "")
	ss = strings.Split(ns, "-")
	ns = strings.Join(ss, "")
	rs := []byte(ns)
	res := make([]byte, 0)
	l := len(ns)
	i := 0
	for l > 3 {
		if l != 4 {
			res = append(res, rs[i:i+3]...)
			res = append(res, '-')
			i += 3
			l -= 3
		} else {
			res = append(res, rs[i:i+2]...)
			res = append(res, '-')
			i += 2
			l -= 2
		}
	}
	if l > 0 {
		res = append(res, rs[i:i+l]...)
		l = 0
	} else {
		res = res[:len(res)-1]
	}
	return string(res)
}

type Cc struct {
	val  int
	step int
}

func minimumOperations(nums []int, start int, goal int) int {
	mp := make(map[int]int, 0)
	queue := make([]Cc, 0)
	queue = append(queue, Cc{val: start, step: 0})
	mp[start] = 0
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		for i := 0; i < len(nums); i++ {
			x := t.val ^ nums[i]
			if x == goal {
				return t.step + 1
			} else {
				if x >= 0 && x <= 1000 {
					if value, ok := mp[x]; !ok || value == 0 {
						queue = append(queue, Cc{
							val:  x,
							step: t.step + 1,
						})
						mp[x] = t.step + 1
					}
				}
			}
			x = t.val - nums[i]
			if x == goal {
				return t.step + 1
			} else {
				if x >= 0 && x <= 1000 {
					if value, ok := mp[x]; !ok || value == 0 {
						queue = append(queue, Cc{
							val:  x,
							step: t.step + 1,
						})
						mp[x] = t.step + 1
					}
				}
			}
			x = t.val + nums[i]
			if x == goal {
				return t.step + 1
			} else {
				if x < 0 || x > 1000 {
					continue
				}
				if value, ok := mp[x]; !ok || value == 0 {
					queue = append(queue, Cc{
						val:  x,
						step: t.step + 1,
					})
					mp[x] = t.step + 1
				}
			}
		}
	}
	return -1
}
func findEvenNumbers(digits []int) []int {
	mp := make(map[int]int, 0)
	for i := 0; i < len(digits); i++ {
		temp := digits[i]
		for j := 0; j < len(digits); j++ {
			if j == i {
				continue
			}
			temp *= 10
			temp += digits[j]
			for k := 0; k < len(digits); k++ {
				if k == i || k == j {
					continue
				}
				temp *= 10
				temp += digits[k]
				if temp%2 == 0 && temp >= 100 && temp <= 999 {
					mp[temp]++
				}
			}
		}
	}
	res := make([]int, 0)
	for key := range mp {
		res = append(res, key)
	}
	return res
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{}
	ret := res
	in := 0
	for l1 != nil && l2 != nil {
		res.Val = l1.Val + l2.Val + in
		//    fmt.Printf("%v, %d\n",res.Val,in)
		if res.Val >= 10 {
			res.Val %= 10
			in = 1
		} else {
			in = 0
		}
		if l1.Next != nil || l2.Next != nil {
			res.Next = &ListNode{}
			res = res.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		res.Val = l1.Val + in
		//    fmt.Printf("%v, %d\n",res.Val,in)
		if res.Val >= 10 {
			res.Val %= 10
			in = 1
		} else {
			in = 0
		}

		if l1.Next != nil {
			res.Next = &ListNode{}
			res = res.Next
		}

		l1 = l1.Next
	}
	for l2 != nil {
		res.Val = l2.Val + in
		// fmt.Printf("l2:%v, %d\n",res.Val,in)
		if res.Val >= 10 {
			res.Val %= 10
			in = 1
		} else {
			in = 0
		}

		if l2.Next != nil {
			// fmt.Printf("f l2:%v, %d\n",res.Val,in)
			res.Next = &ListNode{}
			res = res.Next
		}

		l2 = l2.Next
	}
	if in > 0 {
		//    fmt.Printf("%v, %d\n",res.Val,in)
		res.Next = &ListNode{Val: in}
	}
	return ret
}

//["Bank","withdraw","transfer","transfer","withdraw","withdraw","transfer","deposit","deposit","transfer","transfer","deposit","withdraw","withdraw","withdraw","deposit","transfer","transfer","withdraw","transfer","withdraw","withdraw","withdraw","transfer","withdraw","withdraw","withdraw","transfer","withdraw","withdraw","transfer","deposit","transfer","transfer","deposit","transfer","transfer","transfer","transfer","transfer","transfer","deposit","transfer","withdraw","withdraw","transfer","deposit","withdraw","withdraw","withdraw","transfer","transfer","transfer","deposit","deposit","transfer","deposit","transfer","deposit","withdraw","deposit","deposit","deposit","deposit","transfer","withdraw","deposit","transfer","deposit","withdraw","deposit","withdraw","deposit","deposit","transfer","withdraw","transfer","transfer","withdraw","transfer","withdraw","transfer","deposit","transfer","withdraw","transfer","withdraw","deposit","withdraw","transfer","transfer","withdraw","withdraw","deposit","withdraw","transfer","withdraw","deposit","deposit","deposit","transfer","transfer","transfer","deposit","transfer","withdraw","withdraw","transfer","transfer","deposit","deposit","withdraw","transfer","deposit","deposit","transfer","deposit","deposit","deposit","deposit","withdraw","withdraw","transfer","deposit","transfer","withdraw","deposit","transfer","deposit","withdraw","transfer","withdraw","withdraw","deposit","deposit","transfer","transfer","deposit","withdraw","transfer","withdraw","withdraw","withdraw","transfer","withdraw","deposit","transfer","transfer","withdraw","deposit","deposit","withdraw","deposit","deposit","deposit","deposit","deposit"]
//[[[579,1143,1635,2634,5482,3297,2929,6276,5675,7527,5098,2942,5021,6868,2042,861,1820,7279,1119,8892,6404,3970,2381,8297,4816,4438,7661,8953,5095,3056,778,8597,4212,4963,6180,7206,7078,1277,5274,4951,8258,6131,4255,8297,8971,2959,4404,3890,689,9703,7300,8793,7960,3633,3112,3165,7050,3288,6162,9963,8585,4726,3176,7475,7227,7434,1384,1382,660,6489,1271,8640,4078,7255,9654,9369,6794,6229,7071,1385,7125,8458,3550,592,68,562,7312,7906,4098,4855,1931,1549,5464,7334,5875,5730,2232,1775,8297,8429,7520,6505,7638,6242,3005,6369,5981,870,9937,1056,2042,8185,1806,5585,2096,4233,9198,2549,9070,7604,64,3415,4289,7863,9015,9781,1433,4041,4942,5870,1900,2297,3473,8532,2392,9287,4968,9445,8735,6616,1702,3094,2638,512,1113,1077,5595,8787,221,4147,4877,7851,3267,5194,2150,7845,6253,9183,266,739,1017,8976,869,3638,588,150,6376,8318,1349,4944,1258,3987,5614,4984,4660,8996,9110,4522,6182,3122,4689,6609,3223,9604,2122,5132,1480,7423,6388,1816,3154,2249,6004,6200,6912,777,7846,5966,8576,5559,5648,5314,5960,669,6748,2316,3120,6319,1857,5999,5387,786,7063,3894,6986,2341,1830,6965,3690,2290,8256,4833,1760,5205,2306,1867,2371,5528,6407,6538,4614,2277,3274,6144,7765,5537,8428,2536,6666,5619,226,2877,4592,2252,4855,5394,4580,5916,9990,5685,2375,5836,4726,9866,8692,9112,8286,1216,3286,2260,8379,9402,1322,6215,1356,4029,362,8054,1060,3991,7627,6379,6420,6904,8198,9224,2035,2058,7833,5993,5266,2469,6890,8719,9771,3843,1810,280,1930,9350,1866,6500,5944,396,4460,8483,6272,1554,8695,6308,4480,6000,939,3633,1551,6465,8258,2484,4578,4219,4627,312,9571,6208,154,6852,7278,9652,6416,1466,5651,8157,4730,113,2360,3467,1163,5126,5704,8373,9182,8434,3737,8720,1848,4690,8738,7741,3251,3417,2618,2105,7959,831,7460,7379,2842,3431,4756,6682,5532,8583,1318,2975,8314,4848,907,1569,1949,6728,5278,3098,1327,5994,3092,2025,8196,117,588,5002,9958,8894,3027,2043,6268,9415,5804,9218,8538,9292,4706,3266,9947,3124,5114,4286,1165,5333,9966,4105,1308,2609,7902,3975,9660,5442,4902,4091,6107,5872,8439,7119,6080,8941,860,5524,3875,8600,3664,8427,8379,6476,150,296]],[2421,7151],[151,430,2869],[135,233,715],[1,7889],[75,797],[145,101,8692],[124,8964],[18,1334],[140,287,9090],[161,384,1008],[188,9799],[211,1195],[104,8497],[2948,4080],[129,3741],[321,171,3091],[107,228,1793],[190,701],[332,44,3946],[157,8842],[379,6085],[35,6819],[412,208,309],[2206,9469],[222,1407],[58,1652],[3812,183,1631],[1179,3897],[217,7195],[322,255,4812],[13,3028],[360,322,4090],[41,2975,3798],[117,6158],[116,406,7600],[133,249,4473],[358,307,5807],[217,173,2543],[2,221,5791],[111,204,1454],[219,4829],[179,69,9657],[217,2709],[111,5986],[27,186,9549],[37,4698],[288,2307],[238,6741],[211,6991],[405,64,1483],[1772,222,8026],[192,242,7974],[223,6953],[226,8329],[144,313,2477],[21,1618],[175,68,561],[109,8932],[117,8759],[371,8321],[99,5799],[237,5727],[581,557],[206,161,7402],[111,9363],[371,8775],[365,61,2158],[27,6235],[431,2367],[378,3074],[40,8175],[32,9029],[234,9786],[228,342,4565],[265,2902],[27,241,487],[267,288,8818],[117,6451],[388,279,6601],[349,3789],[283,230,6163],[272,3961],[2958,391,1567],[218,1147],[78,3839,5974],[42,5504],[394,8252],[1365,6333],[87,215,9328],[284,335,4256],[367,2575],[177,5931],[3,8089],[302,6595],[315,361,9208],[3601,7531],[297,782],[767,4355],[104,5557],[227,2681,2369],[149,249,1675],[347,328,1345],[63,7658],[396,297,4770],[386,1448],[2349,7436],[401,2914,4072],[375,2267,9574],[30,3361],[159,935],[209,4745],[3506,362,6961],[368,8707],[327,9819],[277,56,4967],[424,4686],[14,6890],[2567,2857],[286,8185],[307,3484],[151,4524],[206,176,8600],[107,2678],[180,362,9345],[392,6982],[338,6612],[149,30,4531],[299,8293],[157,3055],[402,82,6902],[2363,675],[752,2698],[275,9801],[68,9172],[282,145,7123],[339,213,4217],[182,1334],[43,2310],[29,278,5031],[1511,8965],[1590,250],[1631,7151],[2044,306,1438],[136,4068],[206,4672],[277,62,7349],[294,219,9601],[311,4784],[212,1881],[2638,4787],[3388,5228],[55,3931],[256,9470],[129,9951],[293,4834],[114,6020]]
//[null,false,false,true,false,true,false,true,true,false,true,true,true,false,false,true,true,true,true,true,false,true,false,true,false,true,true,false,false,false,true,true,true,false,true,false,false,false,false,false,true,true,false,false,false,false,true,false,false,false,false,false,false,true,true,false,true,true,true,true,true,true,true,false,false,false,true,true,true,false,true,false,true,true,true,false,true,false,true,false,true,true,true,false,true,false,true,true,false,false,true,true,true,true,false,false,false,true,false,true,false,false,true,true,true,true,false,false,false,true,true,false,false,true,true,false,false,true,false,true,true,false,false,true,false,false,true,false,true,true,true,false,false,true,true,false,false,true,true,true,false,false,false,false,true,true,false,false,false,true,false,false,true,true,true,true,true]
//[null,false,false,true,false,true,false,true,true,false,true,true,true,false,false,true,true,true,true,true,false,true,false,true,false,true,true,false,false,false,true,true,true,false,true,false,false,false,false,false,true,true,false,false,false,false,true,false,false,false,false,false,false,true,true,false,true,true,true,true,true,true,true,false,false,false,true,true,true,false,true,false,true,true,true,false,true,false,true,false,true,true,true,false,true,false,true,true,false,false,true,true,true,true,false,false,false,true,false,true,false,false,true,true,true,true,false,false,false,true,true,false,false,true,true,false,false,true,false,true,true,true,false,true,false,false,true,false,true,true,true,false,false,true,true,false,false,true,true,true,false,false,false,false,true,true,false,false,false,true,false,false,true,true,true,true,true]
func main() {
	fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
	//mp :=make([]map[int]int,2)
	//res:=0
	//for _, _ = range mp {
	//	res++
	//}
	//return res
	nums := []int{1, 3}
	start := 6
	goal := 4
	//fmt.Println(minimumOperations(nums, start, goal))
	nums = []int{2, 4, 12}
	start = 2
	goal = 12
	nums = []int{1}
	start = 0
	goal = 3
	nums = []int{2, 8, 16}
	start = 0
	goal = 1
	nums = []int{77, 74, 20}
	start = 79
	goal = 92
	fmt.Println(minimumOperations(nums, start, goal))
	//L := &ListNode{
	//	Val: 3,
	//}
	//L.Next = &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//fmt.Println(nodesBetweenCriticalPoints(L))
	//head = [5,3,1,2,5,1,2]
	//L2 := &ListNode{
	//	Val: 5,
	//}
	//L2.Next = &ListNode{
	//	Val: 3,
	//}
	//L2.Next.Next = &ListNode{
	//	Val: 1,
	//}
	//L2.Next.Next.Next = &ListNode{
	//	Val: 2,
	//}
	//L2.Next.Next.Next.Next = &ListNode{
	//	Val: 5,
	//}
	//L2.Next.Next.Next.Next.Next = &ListNode{
	//	Val: 1,
	//}
	//L2.Next.Next.Next.Next.Next.Next = &ListNode{
	//	Val: 2,
	//}
	//fmt.Println(nodesBetweenCriticalPoints(L2))
	//s := "**|**|***|"
	//queries := [][]int{{2, 5}, {5, 9}}
	////fmt.Println(platesBetweenCandles(s, queries))
	//s = "***|**|*****|**||**|*"
	//queries = [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}
	//fmt.Println(platesBetweenCandles(s, queries))
	//parents := []int{-1, 2, 0, 2, 0}
	//fmt.Println(countHighestScoreNodes(parents))
	//fmt.Println(nextBeautifulNumber(1))
	//fmt.Println(countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))
	//fmt.Println(countValidWords(" 62   nvtk0wr4f  8 qt3r! w1ph 1l ,e0d 0n 2v 7c.  n06huu2n9 s9   ui4 nsr!d7olr  q-, vqdo!btpmtmui.bb83lf g .!v9-lg 2fyoykex uy5a 8v whvu8 .y sc5 -0n4 zo pfgju 5u 4 3x,3!wl  fv4   s  aig cf j1 a i  8m5o1  !u n!.1tz87d3 .9    n a3  .xb1p9f  b1i a j8s2 cugf l494cx1! hisceovf3 8d93 sg 4r.f1z9w   4- cb r97jo hln3s h2 o .  8dx08as7l!mcmc isa49afk i1 fk,s e !1 ln rt2vhu 4ks4zq c w  o- 6  5!.n8ten0 6mk 2k2y3e335,yj  h p3 5 -0  5g1c  tr49, ,qp9 -v p  7p4v110926wwr h x wklq u zo 16. !8  u63n0c l3 yckifu 1cgz t.i   lh w xa l,jt   hpi ng-gvtk8 9 j u9qfcd!2  kyu42v dmv.cst6i5fo rxhw4wvp2 1 okc8!  z aribcam0  cp-zp,!e x  agj-gb3 !om3934 k vnuo056h g7 t-6j! 8w8fncebuj-lq    inzqhw v39,  f e 9. 50 , ru3r  mbuab  6  wz dw79.av2xp . gbmy gc s6pi pra4fo9fwq k   j-ppy -3vpf   o k4hy3 -!..5s ,2 k5 j p38dtd   !i   b!fgj,nx qgif "))
	//edge := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
	//fmt.Println(secondMinimum(5, edge, 3, 5))
	//nums := []int{3, 1}
	//fmt.Println(countMaxOrSubsets(nums))
	//nums = []int{2, 2, 2}
	//fmt.Println(countMaxOrSubsets(nums))
	//nums = []int{3, 2, 1, 5}
	//fmt.Println(countMaxOrSubsets(nums))
	//balance := []int64{10, 100, 20, 50, 30}
	//bank := Constructor(balance)
	//fmt.Println(bank.Withdraw(3, 10)) // 返回 true ，账户 3 的余额是 $20 ，所以可以取款 $10 。
	//// 账户 3 余额为 $20 - $10 = $10 。
	//fmt.Println(bank.Transfer(5, 1, 20)) // 返回 true ，账户 5 的余额是 $30 ，所以可以转账 $20 。
	//// 账户 5 的余额为 $30 - $20 = $10 ，账户 1 的余额为 $10 + $20 = $30 。
	//fmt.Println(bank.Deposit(5, 20)) // 返回 true ，可以向账户 5 存款 $20 。
	//// 账户 5 的余额为 $10 + $20 = $30 。
	//fmt.Println(bank.Transfer(3, 4, 15)) // 返回 false ，账户 3 的当前余额是 $10 。
	//// 所以无法转账 $15 。
	//fmt.Println(bank.Withdraw(10, 50)) // 返回 false ，交易无效，因为账户 10 并不存在。

	//s := "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
	//fmt.Println(areNumbersAscending(s))
	//s = "hello world 5 x 5"
	//fmt.Println(areNumbersAscending(s))
	//s = "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"
	//fmt.Println(areNumbersAscending(s))
	//s = "4 5 11 26"
	//fmt.Println(areNumbersAscending(s))
	//s := "AAA"
	//fmt.Println(winnerOfGame(s))
	//s = "BBBAAAABB"
	//fmt.Println(winnerOfGame(s))
	//nums := []int{2, 7, 11, 15}
	//target := 9
	//fmt.Println(twoSum(nums, target))
	//nums = []int{3, 2, 4}
	//target = 6
	//fmt.Println(twoSum(nums, target))
	//nums = []int{3, 3}
	//target = 6
	//fmt.Println(twoSum(nums, target))
}
