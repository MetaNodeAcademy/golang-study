package main

import "fmt"

/*
基本数据类型
*/
func main() {
	fmt.Println("Hello World")
	//整型 int int8 int16 int32 int64
	//无符号 int uint uint8 uint16 uint32 uint64
	var a int = 10
	var b int8 = 127
	fmt.Println(a, b)

	//浮点型 float32 float64
	var c float32 = 10.2
	var d float64 = 100.3
	fmt.Println(c, d)

	//布尔型 bool
	var e bool = true
	var f bool = false
	fmt.Println(e, f)

	//byte 类型
	var g byte = 'a'
	fmt.Println(g)
	//string 转bytes
	var ss string = "hello world"
	var by []byte = []byte(ss)
	fmt.Println(by)
	//bytes 转string
	var bys []byte = []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	var str string = string(bys)
	fmt.Println(str)

	var bbb []byte = []byte(ss)
	var ccc []rune = []rune(ss)
	fmt.Println("string 的 length:", len(ss))
	fmt.Println("bytes 的 length:", len(bbb))
	fmt.Println("rune 的 length:", len(ccc))

	//rune int32的内置别名
	var r1 rune = 'q'
	var r2 rune = '国'
	fmt.Println(r1, r2)
	//字符串可以直接转换成[]rune
	var sss string = "haohaohao"
	var rrr []rune = []rune(sss)
	fmt.Println(rrr)
	//环境测试

}
