package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	a := 15
	b := 10
	for a >= b { //while
		b++
		fmt.Println(b)
	}

	for a < 20 {
		fmt.Printf("a=%d\n", a)
		a++
		if a > 17 {
			break //跳出整个循环
		}
	}
	fmt.Println("--------------------------")
	a = 15
	for a < 20 {
		a++
		if a == 18 {
			continue
		}
		fmt.Printf("a=%d\n", a)
	}
	fmt.Println("-------------------------")

	a = 10
LOOP:
	for a < 15 {
		a++
		if a == 12 {
			goto LOOP
		}
		fmt.Printf("a=%d\n", a)
	}
}
