package main

import "fmt"

//常量
func constDemo() {
	const s string = "Hello"
	const a, b = 3, 4
	fmt.Println(s, a, b)

	const (
		s1  = "Go"
		MAX = 10
	)
	//s1 = "Golang" 常量的值不能盖面

}

//枚举
func enumDemo() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Staurday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Staurday, Sunday)
}

//类型定义跟类型别名
func typeDefAlias() {
	type MyInt1 int   //基于Int类型定义的一个新的数据类型
	type MyInt2 = int //Myint2跟int完全一样

	var i int = 1
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func main() {
	constDemo()
	enumDemo()
	typeDefAlias()
}
