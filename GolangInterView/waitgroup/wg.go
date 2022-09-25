package main

import (
	"sync"
	//"time"
)

const N = 10

func main() {
	var wg = sync.WaitGroup{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			//wg.Add(1)
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

}
