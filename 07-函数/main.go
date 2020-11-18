package main

import (
	"fmt"
	"math"
	"os"
)

func operate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("not support oprate: %s\n", op)
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func compute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

type greeting func(name string) string

func say(g greeting, word string) {
	fmt.Println(g(word))
}

func china(name string) string {
	return "你好，" + name
}

func main() {
	a, b := 10, 5
	if result, err := operate(a, b, "x"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	fmt.Println(swap(a, b))

	if file, err := os.Open("abc.txt"); err != nil {
		fmt.Println("打开文件有误")
	} else {
		fmt.Println(file)
	}

	fmt.Println(compute(powInt, 3, 4))

	say(china, "world")
}
