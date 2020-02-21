package main

import "fmt"

func main() {
	var str string = "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)

	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	for i := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

}
