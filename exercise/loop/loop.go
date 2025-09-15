package main

import (
	"fmt"
)

func main() {
	person := [3]string{"张三", "李四", "王五"}
	fmt.Printf("len = %d, cap = %d, array = %v\n", len(person), cap(person), person)

	for key, value := range person {
		fmt.Printf("当前key:%v,当前value:%v\n", key, value)

	}

	for i := range person {
		fmt.Printf("当前index:%v,当前value:%v\n", i, person[i])
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("当前index:%v,当前value:%v\n", i, person[i])
	}

	for _, value := range person {
		fmt.Printf("当前value:%v\n", value)
	}

	//循环slice
	person1 := []string{"张三", "李四", "王五"}
	for key, value := range person1 {
		fmt.Printf("当前key:%v,当前value:%v\n", key, value)

	}
	//循环map
	var person2 map[int]string = make(map[int]string)
	person2[555] = "丁俊晖"
	person2[666] = "墨菲"
	person2[777] = "傅家俊"
	for key, value := range person2 {
		fmt.Printf("当前key:%v,当前value:%v\n", key, value)
	}

	for key := range person2 {
		fmt.Printf("当前key为:%v，当前value为:%v\n", key, person2[key])
	}

	person3 := map[string]string{
		"name":    "张三",
		"age":     "18",
		"sex":     "男",
		"address": "上海",
	}

	for key, value := range person3 {
		fmt.Printf("当前key为:%v，当前value为:%v\n", key, value)
	}

	//break
	for i := 1; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("当前i为:%v\n", i)
	}
	//continue
	for i := range 5 {
		if i == 3 {
			continue
		}
		fmt.Printf("当前i为:%v\n", i)
	}
	//goto
	for i := 1; i < 5; i++ {
		if i == 3 {
			goto myLogic
		}
		fmt.Printf("当前i为:%v\n", i)
	}
myLogic:
	fmt.Println("当前i为3")

	//switch  不需要使用break，当匹配时自动结束
	m := 4
	switch m {
	case 1:
		fmt.Println("当前m为1")
	case 2:
		fmt.Println("当前m为2")
	case 3:
		fmt.Println("当前m为3")
	case 4:
		fmt.Println("当前m为4")
	case 5:
		fmt.Println("当前m为5")
	default:
		fmt.Println("当前m为6")
	}

}
