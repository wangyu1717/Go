package main

import "fmt"

//全局变量
var (
	d1 = 7
	d2 = "golang"
)

func varIntValue() {
	var a, b int = 3, 4
	fmt.Println(a, b)

	var s string = "azu"
	fmt.Println(s)

	var ss = "azu"
	fmt.Println(ss)

	var s1, s2, s3 = 1, "azu", true
	fmt.Println(s1, s2, s3)
}

func main() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
	varIntValue()

	// :=只能在函数内部
	s4, s5, s6 := 11, 12, 13
	fmt.Println(s4, s5, s6)
	fmt.Println(d1, d2)
}
