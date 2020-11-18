package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[string]int //m==nil
	fmt.Println(m, m == nil)
	m = make(map[string]int)
	m["a"] = 10
	fmt.Println(m)

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println(m1, m1["a"], m1["b"])

	if val, key := m1["e"]; key {
		fmt.Println(val)
	} else {
		fmt.Println("key e not map m1 ")
	}
	fmt.Println("=========================")

	for key, val := range m1 {
		fmt.Println(key, val)
	}
	fmt.Println("=========================")

	//排序
	var keys []string
	for key, _ := range m1 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for _, val := range keys {
		fmt.Println(val, m1[val])
	}
}
