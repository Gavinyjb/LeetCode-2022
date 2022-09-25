package main

import (
	"fmt"
	"sync"
)

//func main() {
//	letter, number := make(chan bool), make(chan bool)
//	wait := sync.WaitGroup{}
//
//	go func() {
//		i := 1
//		for {
//			select {
//			case <-number:
//				fmt.Print(i)
//				i++
//				fmt.Print(i)
//				i++
//				letter <- true
//			}
//		}
//	}()
//	wait.Add(1)
//	go func(wait *sync.WaitGroup) {
//		i := 'A'
//		for {
//			select {
//			case <-letter:
//				if i >= 'Z' {
//					wait.Done()
//					return
//				}
//
//				fmt.Print(string(i))
//				i++
//				fmt.Print(string(i))
//				i++
//				number <- true
//			}
//
//		}
//	}(&wait)
//	number <- true
//	wait.Wait()
//}

func main() {
	wg := sync.WaitGroup{}
	numFlag, strFlag := make(chan bool), make(chan bool)
	wg.Add(2)
	go func() {
		i := 0
		for {
			select {
			case _, open := <-numFlag:
				if !open {
					wg.Done()
					return
				}
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				strFlag <- true
			}
		}
	}()
	go func() {
		i := 'A'
		for {
			if i >= 'Z' {
				close(numFlag)
				wg.Done()
				return
			}
			select {
			case <-strFlag:
				fmt.Println(string(i))
				i++
				fmt.Println(string(i))
				i++
				numFlag <- true
			}
		}
	}()
	numFlag <- true
	wg.Wait()
}
