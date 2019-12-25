package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func main() {
	res := somar(1, 2)
	fmt.Println("1 + 2 = ", res)
}
