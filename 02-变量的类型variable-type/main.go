package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

func main() {
	//Go语言中不允许隐式类型转换，只能强制转换
	var a int = 1
	var b int32
	b = int32(a)
	fmt.Println(b)

	var a1, a2 = 3, 4
	var c int
	temp := a1*a1 + a2*a2
	c = int(math.Sqrt(float64(temp)))
	fmt.Println(c)
	//强制类型转换

	//amd64
	cpuArch := runtime.GOARCH
	intSize := strconv.IntSize
	fmt.Println(cpuArch, intSize)

	//浮点数
	var f1 float64
	var f2 float64
	fmt.Printf("%f %f\n", f1, f2)

	var bt byte = 2   //uint8
	var ru rune = '中' //uint32
	fmt.Printf("%T %T\n", bt, ru)

	p := 1
	ptr := &p
	ptrptr := &ptr
	fmt.Println("%T %T %T\n", p, ptr, ptrptr)

}
