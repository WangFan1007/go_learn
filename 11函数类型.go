package main

import "fmt"

func test6() {
	fmt.Println("你好")
}

type FUNC_TYPE func()

func test7(a int, b int) {
	fmt.Println(a + b)
}

type FUNCT2 func(int, int)

func main11() {
	var f FUNC_TYPE
	f = test6

	f()

	f2 := test7
	f2(2, 4)

	fmt.Printf("%T", f2)
}
