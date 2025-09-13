package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var map1 map[int]string = make(map[int]string)
	fmt.Printf("当前map:%v\n", map1)

	map1[0] = "张三"
	map1[1] = "李四"
	map1[2] = "王五"
	fmt.Println(map1)

	map2 := map[int]string{0: "张三1", 1: "李四1", 2: "王五1"}
	fmt.Printf("当前map:%v\n", map2)

	map3 := make(map[string]string)
	map3["a"] = "希金斯"
	map3["b"] = "奥沙利文"
	map3["c"] = "特鲁姆普"

	fmt.Printf("当前map:%v\n", map3)

	//生成json

	var map4 = make(map[string]interface{})
	map4["code"] = 200
	map4["msg"] = "成功"
	map4["data"] = map[string]interface{}{
		"name":  "张三",
		"age":   18,
		"sex":   "男",
		"hobby": []string{"吃饭", "睡觉", "打豆豆"},
	}
	fmt.Printf("当前map:%v\n", map4)
	//序列化
	json1, error := json.Marshal(map4)
	if error != nil {
		fmt.Printf("json序列化失败")
	}
	fmt.Printf("序列化后的json:%v\n", string(json1))
	//反序列化
	jsonStr := string(json1)
	errors := json.Unmarshal([]byte(jsonStr), &map4)
	if errors != nil {
		fmt.Printf("json反序列化失败")
	}
	fmt.Printf("反序列化后的map:%v\n", map4)

	//编辑&删除
	var map5 = map[int]string{1: "张三", 2: "李四", 3: "赵六", 4: "王五"}
	fmt.Printf("当前map:%v\n", map5)
	map5[2] = "宾汉姆"
	map5[3] = "罗伯逊"
	fmt.Printf("修改后的map:%v\n", map5)
	delete(map5, 4)
	fmt.Printf("删除后的map:%v\n", map5)

}
