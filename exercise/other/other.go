package main

import (
	"fmt"
)

func main() {
	fmt.Println("这个文件随便写点")

	a := 2
	b := &a
	//*p取出p对应的内存地址中的值，也就是a的值
	fmt.Printf("b的类型：%T,b的值：%v\n,b对应的内存地址中的值为：%v\n", b, b, *b)

	*b = 88
	fmt.Printf("a的值为：%v\n", a) //这里相当于修改了a的值

	var aa = 1

	var hh *int = &aa
	ff := &aa
	fmt.Printf("ff的引用地址为：%v\n", ff)
	fmt.Printf("pp的引用地址为：%v\n", hh)

	// fn1(aa)
	// fmt.Printf("aa的值为：%v\n", aa)

	// fmt.Printf("aa的引用地址为：%v\n", &aa)

	// fn2(&aa)
	// fmt.Printf("aa的值为：%v\n", aa)

}

func fn1(a int) {
	a = 10
}
func fn2(a *int) {
	*a = 20
}
