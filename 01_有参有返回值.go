package main

import "fmt"

func main() {
	max, min := maxAndMin(10, 20)
	fmt.Printf("max=%v,min=%v", max, min)
}

func maxAndMin(a int, b int) (max int, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}
