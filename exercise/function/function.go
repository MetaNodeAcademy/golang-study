package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var functionA func(int) int

// 声明闭包
// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}

func main() {
	aaa := A{1}

	functionA = aaa.add
	// 声明一个闭包并直接执行
	// 此闭包返回值是另外一个闭包（带参闭包）
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is a anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "test")
	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

	fmt.Println("a.i = ", aaa.i)
	fmt.Println("after call function1, a.i = ", functionA(1))
	fmt.Println("a.i = ", aaa.i)

	//go语言默认是值传递
	var res Result
	res.Code = 200
	res.Message = "操作成功"
	toJson(&res)
	setData(&res)
	toJson(&res)

	//md5
	md5Str := getMd5("123456")
	fmt.Printf("md5:%v\n", md5Str)

	//获取给定时间的时间
	var dataStr string = getCurrentTime()
	fmt.Printf("时间戳:%v\n", dataStr)

	fmt.Printf("时间戳:%v\n", getTimeInt())

	//生成签名
	params := map[string]interface{}{
		"name":     "张三",
		"age":      18,
		"password": "123456",
	}
	fmt.Printf("签名:%v\n", getSign(params))

}

func getSign(para map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range para {
		key = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], para[key[i]])
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], para[key[i]])
		}
	}
	var secret = "12333"

	sign := getMd5(getMd5(str) + getMd5(secret))
	return sign

}

func toJson(res *Result) {
	json, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("json序列化失败")
	}
	fmt.Printf("序列化后的json:%v\n", string(json))

}

func setData(res *Result) {
	res.Code = 222
	res.Message = "修改成功"
}

func getMd5(str string) (res string) {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getTimeInt() int64 {
	return time.Now().Unix()
}
