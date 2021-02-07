package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int
	p := make(chan int, 100)
	wg := sync.WaitGroup{}
	for i = 0; i < 100; i++ {
		p <- i
	}
	close(p)
	wg.Add(2)
	go func() {
		for {
			if num, ok := <-p; ok == true {
				fmt.Println(num)
			} else {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		for {
			if num, ok := <-p; ok == true {
				fmt.Println(num)
			} else {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()

}
