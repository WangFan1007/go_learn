package main

import "fmt"

func main14() {
	arr := [10]int{10, 7, 4, 2, 9, 1, 3, 5, 8, 6}

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
