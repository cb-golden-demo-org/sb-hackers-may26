package main

import "fmt"

// add returns the sum of two integers. This file is deliberately free of
// anything a security scanner would flag (PE UAT test 14a, clean run).
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(2, 3))
}
