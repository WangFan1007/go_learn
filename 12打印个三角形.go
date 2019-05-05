package main

import (
	"fmt"
	"time"
)

func printTriangle() {
	space := "                            "
	for i := 0; i < 10; i++ {
		s := ""
		space = space[i/2:]
		for j := 0; j < i; j++ {
			s += "*"
		}
		fmt.Printf("%s%s%s\n", s, space, s)
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
}

func main() {
	printTriangle()
}
