package main

import "fmt"

func Ref(a int) {
	a = 100
	fmt.Printf("in func inner: %d\n", a)
}

func Ptr(a *int) {
	*a = 100
	fmt.Printf("in func inner: %d\n", *a)
}

func main() {
	p := 10
	Ref(p)
	fmt.Println(p)

	Ptr(&p)
	fmt.Println(p)
}
