package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	// 初始化一个context
	parent := context.Background()
	// 生成一个取消类
	ctx, cancel := context.WithCancel(parent)
	runTime := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine Done")
				return
			default:
				fmt.Printf("Goroutine Running Times : %d \n", runTime)
				runTime = runTime + 1
			}
			if runTime > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)

	wg.Wait()
}
