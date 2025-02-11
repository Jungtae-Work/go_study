package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	expiry := time.Date(2024, 10, 5, 21, 4, 0, 0, time.Local) // 2024-10-05 21:00:00

	if expiry.Compare(time.Now()) <= 0 {
		fmt.Println("The Holdem Manager has expired.")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go func(x int) {
			time.Sleep(time.Duration(x) * 100 * time.Millisecond)
			ticker := time.NewTicker(300 * time.Millisecond)
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
	time.Sleep(1 * time.Second)
	fmt.Println("Program End.")
}
