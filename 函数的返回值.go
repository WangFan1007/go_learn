package main

import "fmt"

func test() (a int, b int, c int) {
	a = 10
	b = 20
	c = 30

	return
}

func main7() {
	a, b, c := test()
	fmt.Println("values:", a, b, c)
}
