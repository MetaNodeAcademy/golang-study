package main

import "fmt"

func main() {

	//一维数组
	var arr [5]int
	fmt.Println(arr)

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"张", "三", "言"}
	fmt.Println(arr2)

	arr3 := [...]float32{1.2, 4.5, 5.6, 6.6}
	fmt.Println(arr3)

	arr4 := [5]int{1, 2}
	fmt.Println(arr4)

	arr5 := [8]int{0: 1, 1: 55, 2: 66}
	fmt.Println(arr5)

	//二维数组
	var arr6 [3][2]string = [3][2]string{{"张", "三"}, {"李", "四"}, {"王", "五"}}
	fmt.Println(arr6)

	arr7 := [5][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr7)

	arr8 := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr8)

	var arr9 = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr9)
	fmt.Println(arr9)

	var arr10 = [5]int{1, 2, 3, 4, 5}
	modifyArr1(&arr10)
	fmt.Println(arr10)

	arr11 := [5]int{1, 2, 3, 4, 5}
	var arr12 [5]int = arr11 //同类型同个数的才能赋值
	fmt.Printf("arr12的值为：%v\n", arr12)

}

func modifyArr1(arr *[5]int) {
	arr[0] = 150
}

func modifyArr(arr [5]int) {
	arr[0] = 100
}
