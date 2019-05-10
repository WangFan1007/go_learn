package main

import (
	"fmt"
	"time"
)

func main18() {
	fmt.Println(time.Now())

	ticker := time.NewTicker(time.Second * 1)
	is8Times := make(chan bool)
	go func() {
		i := 0

		for {
			now := <-ticker.C
			fmt.Println(now)

			if i == 8 {
				is8Times <- true
			}
			i++
		}
	}()

	<-is8Times
}
