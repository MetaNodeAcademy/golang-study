package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p1 *int
	var p2 *string

	fmt.Println(p1)
	fmt.Println(p2) //nil

	var num = 2
	var str = "张三"

	p1 = &num
	p2 = &str
	var p3 = &p2

	fmt.Println(p1)
	fmt.Println(p2) //内存地址 0xc0000280a0
	fmt.Println(p3)

	//使用指针访问值
	var p4 *int
	var p5 int = 5

	p4 = &p5
	fmt.Printf("p4的取内存地址：%v\n", &p4)    //&--->取内存地址
	fmt.Printf("p4的内存地址对应的值：%v\n", *p4) //*--->取内存地址对应的值

	//修改指针指向的值
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)

	/*
		uintptr 类型是把内存地址转换成了一个整数，然后对这个整数进行计算后，在把 uintptr 转换成指针，达到指针偏移的效果。
		unsafe.Pointer 是普通指针与 uintptr 之间的桥梁，通过 unsafe.Pointer 实现三者的相互转换
	*/
	sss := "你好啊"
	str1 := uintptr(unsafe.Pointer(&sss))
	fmt.Println(str1)
	str1 += 1

	fmt.Println(str1)

}
