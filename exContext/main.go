package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go func(x int) {
			ticker := time.NewTicker(500 * time.Millisecond)
			cnt := 0

			fmt.Println(x, "GoRoutine Start", cnt)
			for {
				select {
				case <-ctx.Done():
					fmt.Println(x, "GoRoutine End", cnt)
					return
				case t := <-ticker.C:
					cnt++
					fmt.Println("==>", x, cnt, t)
				default:
					time.Sleep(100 * time.Millisecond)
					// fmt.Println(x)
				}
			}
		}(i)
	}

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Program End.")
}
