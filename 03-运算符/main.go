package main

import "fmt"

//算术运算符
func operateDemo() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("a+b=%d\n", c)
	//c := a + b

	c = a - b
	fmt.Printf("a-b=%d\n", c)

	c = a * b
	fmt.Printf("a*b=%d\n", c)

	c = a / b
	fmt.Printf("a/b=%d\n", c)

	c = a % b
	fmt.Printf("a%%b=%d\n", c)

	a++
	fmt.Printf("a++=%d\n", a)

	a = 21

	a--
	fmt.Printf("a--=%d\n", a)

}

//关系运算符
func relationDemo() {
	var a int = 21
	var b int = 10

	if a == b { //false
		fmt.Println("a==b")
	} else {
		fmt.Println("a!=b")
	}

	if a > b { //true
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}

	if a < b { //false
		fmt.Println("a < b")
	} else {
		fmt.Println("a >= b")
	}

}

//逻辑运算符
func logicDemo() {
	var a bool = true
	var b bool = false

	if a && b { //如果a和b都是真     //false
		fmt.Println("a和b都为真")
	} else {
		fmt.Println("a和b有假")
	}

	if a || b { //如果a和b中有一个为真 //true
		fmt.Println("a或者b至少有一个为真")
	} else {
		fmt.Println("a和b都为假")
	}

	if !a { //a为真, !a非真即为假
		fmt.Println("a为假")
	} else {
		fmt.Println("a为真")
	}
}

//位运算符
func byteDemo() {
	var a uint = 60 //60=111100
	fmt.Printf("%b\n", a)
	var b uint = 13 //13=1101
	fmt.Printf("%b\n", b)

	var c uint

	c = a & b
	//c=1100
	fmt.Printf("a&b=%d\n", c)

	c = a | b
	//c =111101
	fmt.Printf("a|b=%d\n", c)

	c = a ^ b
	//c=110001
	fmt.Printf("a^b=%d\n", c)

	c = a << 2 //60 * 4 = 240
	// a << n 左移动n位=乘以2的n次方
	fmt.Printf("a<<2=%d\n", c)

	c = a >> 2 //60 * 4 = 15
	// a >>n 右移n位=除以2的n次方
	fmt.Printf("a>>2=%d\n", c)

}

func main() {
	operateDemo()
	relationDemo()
	logicDemo()
	byteDemo()
}
