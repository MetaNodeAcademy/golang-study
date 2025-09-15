package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	var res Result
	res.Code = 200
	res.Message = "操作成功"
	res.Data = "这是数据"
	fmt.Printf("当前res:%v\n", res)

	//序列化数据
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("json序列化失败")
	}
	res1 := string(result)
	fmt.Printf("序列化后的json:%v\n", res1)
	//反序列化
	var res2 Result
	err1 := json.Unmarshal([]byte(result), &res2)
	if err1 != nil {
		fmt.Printf("json反序列化失败")
	}
	fmt.Printf("反序列化后的res:%v\n", res2)

	var p1 Person = Person{name: "张三", age: 12, sex: "男", hobby: []string{"打台球", "写代码", "游泳"}}
	fmt.Printf("当前p1:%v\n", p1)

	p2 := Person{name: "李四", age: 33, sex: "男", hobby: []string{"跑步", "打球", "游泳"}}
	fmt.Printf("当前p2:%v\n", p2)

	//匿名结构体
	p4 := struct {
		name   string
		Height int
		Wight  float64
	}{name: "王五",
		Height: 180,
		Wight:  180.0,
	}
	fmt.Printf("当前p4:%v\n", p4)

	//修改数据
	toJson(&res)
	fmt.Printf("修改前的res:%v\n", res)
	setData(&res)
	toJson(&res)
	fmt.Printf("修改后的res:%v\n", res)

}

func setData(res *Result) {
	res.Code = 222
	res.Message = "修改成功"
	res.Data = "这是修改后的数据"
}

func toJson(res *Result) {
	json, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("json序列化失败")
	}
	fmt.Printf("序列化后的json:%v\n", string(json))
}
