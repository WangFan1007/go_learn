package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var value int

var mutext = sync.RWMutex{}

func write(idx int) {
	rand.Seed(time.Now().UnixNano())
	for {
		mutext.Lock()
		num := rand.Intn(1000)
		value = num
		fmt.Printf("写入%d go程%d\n", num, idx)
		time.Sleep(time.Second * 1)
		mutext.Unlock()
	}
}

func read(idx int) {
	for {
		mutext.RLock()
		num := value
		fmt.Printf("读取到 %d go程%d\n", num, idx)
		mutext.RUnlock()
	}
}

func main20() {
	for i := 0; i < 5; i++ {
		go write(i)
	}
	for i := 0; i < 5; i++ {
		go read(i)
	}

	for {

	}
}
