package main

import (
	"fmt"
	"os"
)

func main() {
	for _, ars := range os.Args[1:] {
		fmt.Println(ars)
	}
}
