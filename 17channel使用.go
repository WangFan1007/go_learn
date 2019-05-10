package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 100)
	}

}

func test1() {
	printer("hello")
	channel <- 8
}

func test2() {
	<-channel
	printer(" world")
}

func main17() {
	go test1()
	go test2()

	for {

	}
}
