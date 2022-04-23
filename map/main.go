package main

import "fmt"

func main() {
	// var m map[string]int
	// m["key"] = 1
	// fmt.Println(m) // 发生运行时异常：panic: assignment to entry in nil map
	// n := map[string]int{}
	// n["key"]=1
	// fmt.Println(n)

	// comma ok
	m := make(map[string]int)
	v,ok := m["key1"]
	if !ok{
		fmt.Println("查无此key")
	}else {
		fmt.Println(v)
	}
	
}