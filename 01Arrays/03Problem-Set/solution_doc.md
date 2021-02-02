# 剑指Offer数组类型题解

## 面试题03 数组中重复的数字

链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

<img src="https://gitee.com/oluoluo/typoraImage/raw/master/img/image-20210130210609566.png" alt="image-20210130210609566" style="zoom:90%;" />

第一种方法很自然能想到，使用标记数组或者集合，时间复杂度是`O(N)`，空间复杂度是`O(N)`，这没什么好说的。如果加一个要求：**要求空间复杂度是`O(1)`呢**？

网上看到的第一种方法是原地交换，利用题目中的条件，在一个长度为`n`的数组`nums`里的所有数字都在0 ~ n-1之间，那么数组的索引与数组的值之间存在一对多的映射关系。这就是数学上的“鸽巢原理”。原地交换的思想就是：如果当前位置`i`的元素`val`不等于`i`（鸽子不在巢中），判断位置`val`的元素是否为`val`，是的话说明该数字重复直接返回，不是的话就将两个位置的元素交换；如果当前位置`i`的元素`val`等于`i`，继续访问数组下一个元素。

<img src="https://gitee.com/oluoluo/typoraImage/raw/master/img/image-20210130213241443.png" alt="image-20210130213241443" style="zoom:67%;" />

实现：

```go
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		} else {
			val := nums[i]
			if nums[val] == val {
				return val
			} else {
				nums[i], nums[val] = nums[val], nums[i]
			}
		}
	}
	return -1
}
```

## 面试题04 二维数组中的查找

链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

题目描述：在一个`n * m`的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下的顺序排序，输入这样一个二维数组和一个整数，判断数组中是否含有该整数。

暴力的方法时间复杂度太大肯定不行。

一开始想的是借鉴二分查找的思想，在数组中间选择一个数，根据它的大小判断要查找的数字可能出现的区域。如果选择数字小于待查找的数字，则待查找的数字可能出现在如下图a中的阴影区域；同样，如果选择数字大于待查找的数字，则查找的数字可能出现在图b中的阴影区域。并不好处理。

<img src="https://gitee.com/oluoluo/typoraImage/raw/master/img/image-20210131183936097.png" alt="image-20210131183936097" style="zoom:67%;" />

没想出来，参考了一下题解，处理的办法是选择矩阵右上角的数字，如果要查找的数字刚好等于矩阵右上角的数字，则返回；如果要查找的数字小于矩阵右上角的数字，那么往左查找，选择的数字所在列就被剔除了（其中的数字都要大于要查找的数字）；如果要查找的数字大于矩阵右上角的数字，那么往下查找，选择的数字所在的行就被剔除了（选择的数字的左侧所在行的数字都要小于待查找的数字）。实现如下：

```go
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, column := 0, len(matrix[0])-1
	for {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
		if row >= len(matrix) || column < 0 {
			return false
		}
	}
}
```

## 面试题03 替换空格

链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

如果对空间复杂度没有要求的话，直接使用`bytes.Buffer`就好。实现：

```go
func replaceSpace(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
```

## 面试题29 顺时针打印数组

链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

这道题是之前群里的小伙伴在群里分享的一道腾讯面试题，这道题看起来很简单，但写起来也不算是太容易。我的思路是先顺时针访问最外层，再依次访问内层，第一遍写出来的代码在列数为1或者行数为1的时候出现一些数字访问了两遍的问题，所以加上了一个条件判断。实现如下：

```go
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	column, row := len(matrix[0]), len(matrix)
	var retList []int = make([]int, 0)
	ix, iy := 0, 0
	for column > 0 && row > 0 {
		for i := 0; i < column; i++ {
			retList = append(retList, matrix[ix][iy+i])
		}
		for i := 1; i < row; i++ {
			retList = append(retList, matrix[ix+i][iy+column-1])
		}
		if column != 1 && row != 1 {
			for i := 1; i < column; i++ {
				retList = append(retList, matrix[ix+row-1][iy+column-1-i])
			}
			for i := 1; i < row-1; i++ {
				retList = append(retList, matrix[ix+row-1-i][iy])
			}
		}
		ix++
		iy++
		column -= 2
		row -= 2
	}
	return retList
}
```

## 面试题50 第一个只出现一次的字符

链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

题目描述：在字符串s中找出第一个只出现一次的字符，如果没有，返回一个单空格，s只包含小写字母。

最开始的解法比较局限，即遍历一遍字符串，将字符、次数统计、出现顺序都放在一个结构体中，然后对存放结构体的哈希表（数组实现，索引是小写字母对应的0-25）按照结构体中的出现顺序从小到大排序，然后返回第一个次数统计值为1的字符。这种解法因为使用了排序，时间复杂度是`O(nlogn)`。显然不是最优的。相关实现：

```go
func firstUniqChar(s string) byte {
	hashtable := make([]pair, 26)
	oid := 1
	for i := 0; i < len(s); i++ {
		if hashtable[s[i]-'a'].cnt == 0 {
			hashtable[s[i]-'a'].order = oid
			oid++
			hashtable[s[i]-'a'].ch = s[i]
		}
		hashtable[s[i]-'a'].cnt++
	}
	sort.Slice(hashtable, func(i, j int) bool {
		return hashtable[i].order < hashtable[j].order
	})
	for i := 0; i < 26; i++ {
		if hashtable[i].cnt == 1 {
			return hashtable[i].ch
		}
	}
	return ' '
}
```

第二种方法的时间复杂度是`O(n)`，只需要遍历两次字符串就可以了。第一次遍历字符串统计各种字符出现的次数，第二次遍历字符串找到最早出现并且只出现一次的字符。然后使用golang自带的内置数据结构map的话执行时间只超过了10%，如果使用数组实现哈希表的话执行时间超过100%。

```go
// 第二种解法，访问两遍字符串，不使用map执行时间更短
func firstUniqChar1(s string) byte {
	// mp := make(map[byte]int)
	// for i := 0; i < len(s); i++ {
	// 	mp[s[i]] = mp[s[i]] + 1
	// }
	// for i := 0; i < len(s); i++ {
	// 	if mp[s[i]] == 1 {
	// 		return s[i]
	// 	}
	// }
	var hashtable [26]int
	for i := 0; i < len(s); i++ {
		hashtable[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if hashtable[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
```

<img src="https://gitee.com/oluoluo/typoraImage/raw/master/img/image-20210202185424174.png" alt="image-20210202185424174" style="zoom:67%;" />