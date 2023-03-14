package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Printf("i:%v, s[i]:%v\n", i, s[i])
	}
	fmt.Println(s[0:1])
}
