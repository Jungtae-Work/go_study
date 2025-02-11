package main

import (
	"fmt"
	"id-queue/game"
)

func main() {
	que1 := game.NewIdQueue(1, 100)

	for i := 0; i < 10; i++ {
		a := que1.Pop()

		fmt.Println(i, a)
	}
	fmt.Println(&que1)

	que1.Push(15)
	que1.Push(20)
	que1.Push(17)

	fmt.Println(&que1)
	que1.Push(que1.Pop())
	fmt.Println(&que1)
}
