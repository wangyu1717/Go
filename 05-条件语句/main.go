package main

import "fmt"

func main() {
	a := 22
	b := 2
	if a < 20 { //条件表达式 true false     //false
		fmt.Println("a < 20")
	} else if a == 20 { //false
		fmt.Println("a = 20")
	} else {
		fmt.Println("a > 20")

		if b > 5 {
			fmt.Println("b >5")
		} else {
			fmt.Println("b <= 5")
		}
	}

	fmt.Println("------------------------------")

	switch a {
	case 20:
		fmt.Println("a=20")
	case 21:
		fmt.Println("a=21")
	case 22:
		fmt.Println("a=22")
	default:
		fmt.Println("Unknown")
	}
}
