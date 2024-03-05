package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["s"] = 1
	v := m["s"]
	v = 2
	fmt.Println(m["s"], v)
}
