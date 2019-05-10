package main

import (
	"fmt"
	"time"
)

func singing() {
	for i := 0; i < 5; i++ {
		fmt.Println("singing")
		time.Sleep(1)
	}

}

func dancing() {
	for i := 0; i < 10; i++ {
		fmt.Println("dancing")
		time.Sleep(1)
	}
}

func main16() {
	go singing()
	go dancing()
	for {
		time.Sleep(500)
	}
}
