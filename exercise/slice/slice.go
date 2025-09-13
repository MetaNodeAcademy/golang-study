package main

import "fmt"

func main() {
	fmt.Println("切片-动态数组，比数组灵活")

	/*
	 len():获取切片的实际长度
	 cap():获取未扩容状态下的切片容量，若切片扩容，容量会动态调整
	*/

	//声明切片

	var slice []string
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice), cap(slice), slice)

	var slice1 = []int{}
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice1), cap(slice1), slice1)

	var slice2 = [5]string{"李", "王", "张", "刘", "熊"}
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice2), cap(slice2), slice2)

	slice3 := make([]int, 5, 59)
	slice3[0] = 111
	slice3[1] = 222
	slice3[3] = 555
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice3), cap(slice3), slice3)

	//截取切片
	slice4 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice4), cap(slice4), slice4)

	fmt.Printf("截取切片的第一个元素：%d\n", slice4[0])

	fmt.Printf("截取切片的索引为1到4的元素：%v\n", slice4[1:5])

	//切片追加
	var slice5 []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice5), cap(slice5), slice5)
	slice5 = append(slice5, 10) //触发扩容，容量翻倍即cap = 14
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice5), cap(slice5), slice5)
	slice5 = append(slice5, 11, 12, 13)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice5), cap(slice5), slice5)

	//删除切片元素
	var slice6 []int = []int{1, 2, 3, 4, 5, 6, 7}

	//删除头部2个元素

	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice6[2:]), cap(slice6[2:]), slice6[2:])

	slice6 = slice6[2:] //相当于java中的substring(2, slice6.length)
	fmt.Println(slice6)

	//删除尾部两个元素
	slice6 = slice6[:len(slice6)-2]
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(slice6), cap(slice6), slice6)

}
