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

## addTwoNumbers

题意：

将两个倒序使用链表存放的数字相加,输出和(使用链表表现)

题解：

模拟，使用一个进位标记符号，遍历两个链表，如果其中一个为空就，计为0，直至链表为空，进位标记符也为0

时间复杂度：

O(n),max(len(l1),len(l2))

空间复杂度

O(1)，用于计数

## 面试题 01.08. 零矩阵

题意:

若M × N矩阵中某个元素为0，则将其所在的行与列清零。

题解：

使用两个一维数组，分别记录对应的行列是否有0，然后对含有0的行列所有元素清零。

时间复杂度:
O(n*m)
空间复杂度:
O(n+m)

## 1694. 重新格式化电话号码

> 题意:
>- 给一个电话号码字符串,由数字、空格 <kbd>' '</kbd>、和破折号 <kbd>'-'</kbd> 组成。
>- 先去除所有的删除 所有的空格和破折号
>- 从左到右，将数字2个或者3个组成一组，用破折号连接
>- 例子
>- 输入：number = "1-23-45 6"
>- 输出："123-456"

题解：
使用库函数去除空格和破折号

时间复杂度:
O(len(s)) 
空间复杂度:
O(len(s))