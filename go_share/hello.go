package main

import "fmt"

func main() {
	a := [...]int{3, 48, 3423, 1, 3, 435}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9, 1, 1, 1, 1, 1, 1, 1, 1}
	copy(s2[2:4], s1[1:3])
	fmt.Println(s2)

	ch := 'a'
	str := "str"
	fmt.Printf("ch=%T\n", ch)
	fmt.Printf("str=%T", str)

}
