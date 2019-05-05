package main

import (
	"bufio"
	"fmt"
	"os"
)

func main3() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("要输入文件参数")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// fmt.Println(counts)

	for line, n := range counts {
		// if n > 1 {
		fmt.Printf("%d\t%s\n", n, line)
		// }
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
