package main

import (
	"fmt"
	"os"
)

func main2() {
	for _, ars := range os.Args[1:] {
		fmt.Println(ars)
	}
}
