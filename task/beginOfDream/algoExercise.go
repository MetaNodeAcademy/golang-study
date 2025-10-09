package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("开始算法练习")

	/*
		1.只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，
		其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
		结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	*/
	var arr1 = []int{1, 1, 2, 2, 3, 3, 4, 5, 5}
	fetchOnceNumber(arr1)

	/*
		2.判断一个整数是否是回文数
	*/
	var num = -123321
	//将num转换成字符数组并循环
	hwNumber(num)

	/*
		3.给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
		有效字符串需满足：
		1.左括号必须用相同类型的右括号闭合。
		2.左括号必须以正确的顺序闭合。
		3.每个右括号都有一个对应的相同类型的左括号。
	*/

	var s = "()[]{}"
	flag := validStr(s)
	if flag {
		fmt.Println("有效字符串")
	} else {
		fmt.Println("无效字符串")
	}

	/*
		4.最长公共前缀
	*/
	var strs = []string{"123", "12345", "12657575"}
	prefix := longestCommonPrefix(strs)
	if prefix != "" {
		fmt.Println("最长公共前缀为：", prefix)
	} else {
		fmt.Println("没有最长公共前缀")
	}

	/*
		5.给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
		这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
		将大整数加 1，并返回结果的数字数组。
	*/
	digits := []int{1, 7, 3}
	fmt.Println(plusOneArray(digits))

	/*
		6.删除有序数组中的重复项
		给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
		返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
		可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i]
		与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
	*/

	var originArray []int = []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9}
	finalArray := getFinalArray(originArray)
	fmt.Printf("经过去重后的finalArray = %v\n", finalArray)

	/*
			7.两数之和
			给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出
			和为目标值 target  的那 两个 整数，并返回它们的数组下标。
		    你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
	*/

	target := 9
	var originArray2 []int = []int{1, 2, 6, 5, 7, 9}
	finalIndexArray := towSum(originArray2, target)
	if len(finalIndexArray) == 0 {
		fmt.Println("没有满足条件的索引")
	} else {
		fmt.Printf("最终满足条件的索引数组为：%v", finalIndexArray)
	}

}

func towSum(originArray []int, target int) []int {
	for i := 0; i < len(originArray); i++ {
		for j := i + 1; j < len(originArray); j++ {
			if originArray[i]+originArray[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func getFinalArray(originArray []int) []int {
	if len(originArray) == 0 {
		return []int{}
	}
	var finalArray []int
	finalArray = append(finalArray, originArray[0]) // 先添加第一个元素
	for i := 1; i < len(originArray); i++ {         // 从第二个元素开始遍历
		if originArray[i] != originArray[i-1] {
			finalArray = append(finalArray, originArray[i])
		}
	}
	return finalArray
}

func plusOneArray(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前位小于9，直接加1并返回结果
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 如果所有位都是9（如[9,9]），需要在数组最前面添加1
	// 使用append和切片操作在开头插入1
	digits = append([]int{1}, digits...) //digits 切片中的所有元素展开，作为单独的参数传递给 append 函数
	return digits
}

func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	for i := 0; i < len(str[0]); i++ { //以第一个来匹配
		char := str[0][i]
		// 检查其他所有字符串在位置i是否也有相同的字符
		for j := 1; j < len(str); j++ {
			if i >= len(str[j]) || str[j][i] != char {
				// 返回匹配的部分（不包括位置i）
				return str[0][:i]
			}

		}

	}
	return str[0]
}

func validStr(str string) bool {
	if str == "" || len(str) == 0 || len(str)%2 != 0 {
		return false
	}
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, v := range str {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
				return false
			}
			stack = stack[:len(stack)-1] //匹配成功后出栈(移除了切片的最后一个元素)
		}

	}
	return len(stack) == 0
}

func hwNumber(num int) {
	var a1 = []byte(strconv.Itoa(num))
	var str string
	for i := len(a1) - 1; i >= 0; i-- {
		str += string(a1[i])
	}
	if str == string(a1) {
		fmt.Printf("%d 是回文数", num)
	} else {
		fmt.Printf("%d 不是回文数", num)
	}
}

func fetchOnceNumber(arr []int) {
	var temMap = make(map[int]int)

	for _, v := range arr {
		if value, ok := temMap[v]; ok {
			//存在key
			temMap[v] = value + 1
		} else {
			//不存在key
			temMap[v] = 1
		}
	}

	for k, v := range temMap {
		if v == 1 {
			fmt.Printf("找到一个不重复的元素：%d\n", k)
		}
	}
}
