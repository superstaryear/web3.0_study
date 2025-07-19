package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
###
第一题
- **[136. 只出现一次的数字](https://leetcode.cn/problems/single-number/)**：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 `for` 循环遍历数组，结合 `if` 条件判断和 `map` 数据结构来解决，例如通过 `map` 记录每个元素出现的次数，然后再遍历 `map` 找到出现次数为1的元素。
*/
func singleNumber(nums []int) int {
	numsMap := make(map[int]int, len(nums))
	//fmt.Println(m)
	for i := range nums {
		value := nums[i]
		//fmt.Println(value)
		count, exisit := numsMap[value]
		if exisit {
			numsMap[value] = count + 1
		} else {
			numsMap[value] = 1
		}
		//fmt.Printf("index=%d,value=%d\n", key, value)
	}
	var firstKey int
	for key, value := range numsMap {
		if value <= 1 {
			firstKey = key
			break
		}
	}
	return firstKey
}

/*
第二题
题目：判断一个整数是否是回文数
*/
func isPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	//fmt.Println(numStr)
	runes := []rune(numStr)
	//fmt.Println(len(runes))
	var reverse string
	for i := len(runes) - 1; i >= 0; i-- {
		reverse += string(runes[i])
	}
	if numStr == reverse {
		return true
	}
	return false
}

/*
第三题
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func isValid(str string) bool {
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)
	runes := []rune(str)
	for _, char := range runes {
		//fmt.Println(string(char))
		matching, isRight := pairs[char]
		//fmt.Println(string(matching), isRight)
		//如果是右括号
		if isRight {
			if len(stack) == 0 || stack[len(stack)-1] != matching {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			//左括号入栈
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

/*
第四题
最长公共前缀  查找字符串数组中的最长公共前缀
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	fmt.Println(len(strs))
	for i := 0; i < len(strs[0]); i++ {
		currentChar := strs[0][i]
		//fmt.Printf("%c", currentChar)
		//fmt.Println(string(strs[0][i]))
		// 检查其他字符串的相同位置
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够或字符不匹配
			if i >= len(strs[j]) || strs[j][i] != currentChar {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

/*
第五题
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(digits []int) []int {
	var length = len(digits)
	//fmt.Println(length)
	for i := length - 1; i >= 0; i-- {
		//fmt.Printf("%d ", digits[i])
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

/*
第六题
删除有序数组中的重复项
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func removeDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	//慢指针i从0开始，记录不重复元素的位置
	var i = 0
	// 快指针j从1开始，遍历整个数组
	for j := 1; j < len(nums); j++ {
		// 当发现不相同的元素时
		if nums[i] != nums[j] {
			// 将不重复元素移动到i+1位置
			i++
			nums[i] = nums[j]
		}
		// 如果相同，j继续前进，i保持不变

	}
	// 返回去重后的切片
	return nums[:i+1]
}

/*
第七题
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中
*/
func merge(intervals [][]int) [][]int {
	//1.特殊情况处理，空数组或只有一个区间
	if len(intervals) <= 1 {
		return intervals
	}
	//2.按照区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//3.初始化结果集，放入第一个区间
	var merged [][]int = [][]int{intervals[0]}
	//fmt.Println(merged)

	//4.遍历排序后的区间
	for i := 1; i < len(intervals); i++ {
		var last = merged[len(merged)-1]
		var current = intervals[i]
		//5.检查是否有重叠
		if current[0] <= last[1] {
			//有重叠，合并区间（取最大的结束值）
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			//无重叠，直接添加到结果集
			merged = append(merged, current)
		}
	}
	return merged
}

/*
第八题
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func twoSum(nums []int, target int) []int {
	var hashMap = make(map[int]int)
	for i, num := range nums {
		//fmt.Println(i, num)
		complement := target - num
		if j, ok := hashMap[complement]; ok {
			return []int{j, i}
		}
		hashMap[num] = i
	}
	return nil
}
func main() {
	//1、只出现一次的数字
	//var nums = []int{1, 2, 2, 1, 5, 5, 7}
	//var singleNumberResult = singleNumber(nums)
	//fmt.Println(singleNumberResult)

	//var num = 121
	//2、回文数
	//var isPalindromeFlag bool = isPalindrome(num)
	//fmt.Println(isPalindromeFlag)

	//3、判断字符串是否有效
	//var isValidFlag = isValid("{[]}")
	//fmt.Println(isValidFlag)

	//4、查找字符串数组中的最长公共前缀
	//var str = []string{"flower", "flow", "flight"}
	//fmt.Println(longestCommonPrefix(str))

	//5、给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	//var nums []int = []int{9, 9, 9}
	//fmt.Println(plusOne(nums))

	//6、删除有序数组中的重复项
	//var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//fmt.Println(removeDuplicates(nums))

	//7、合并区间
	//var nums = [][]int{{2, 6}, {8, 10}, {15, 18}, {1, 3}}
	//fmt.Println(merge(nums))

	//8、两数之和
	//var nums = []int{2, 11, 15, 7}
	//var target = 9
	//fmt.Println(twoSum(nums, target))

}
