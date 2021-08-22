# leetcode-hub-go

## twoSum

题意：

寻找一个数组里面的两数之和为target的数组下标id

题解：

使用map记录，key为值，value为数组下标

遍历的时候，先寻找，当存在map[target-nums[i]]时，即返回

如果不存在，增加当前值

时间复杂度：

O(n）

空间复杂度

O(n)