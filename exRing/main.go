package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create two rings, r and s, of size 2
	r := ring.New(5)

	// Get the length of the ring
	lr := r.Len()

	// Initialize r with 0s
	for i := 0; i < lr; i++ {
		r.Value = i + 1
		r = r.Next()
	}

	rh := r
	view("Head", rh)
	view("Ring", r)

	//r = r.Move(2)
	//view("Ring+2", r)
	r = r.Prev()
	r.Unlink(lr - 1)

	view("Head", rh)
	view("Ring", r)
}

func view(name string, data *ring.Ring) {
	fmt.Printf("[%s] ", name)
	data.Do(func(p any) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println()
}
