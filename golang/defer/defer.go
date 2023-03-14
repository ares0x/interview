package main

import "fmt"

type num int

func (n num) print() {
	fmt.Println(n)
}

func (n *num) pprint() {
	fmt.Println(*n)
}

func main() {
	//	var n num
	//	defer n.print()
	//	defer n.pprint()
	//	defer func() {
	//		n.print()
	//	}()
	//	defer func() {
	//		n.pprint()
	//	}()
	//	n = 3
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	return

}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
