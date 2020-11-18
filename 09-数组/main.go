package main

import "fmt"

func changeEle(arr *[3]int) {
	arr[0] = 1000
}

func main() {
	var arr1 [5]int //默认值
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3} //初始化并赋值
	fmt.Println(arr2)

	arr3 := [...]int{7, 8, 9}
	fmt.Println(len(arr3), arr3)

	for i, num := range arr3 {
		fmt.Println(i, num)
	}

	//for _, num := range arr3 {
	//	fmt.Println(num)
	//}
	//
	//for i, _ := range arr3 {
	//	fmt.Println(i)
	//}

	//arr5 := [3]int{1, 2, 3}
	//changeEle(arr5)
	//fmt.Println(arr5)

	arr6 := [3]int{1, 2, 3}
	changeEle(&arr6)
	fmt.Println(arr6)

}
