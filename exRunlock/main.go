package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	lockFile := "/tmp/my_program.lock"

	file, err := os.OpenFile(lockFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to create or open lock file:", err)
		return
	}
	defer file.Close()

	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		fmt.Println("Another instance is already running. Exiting.")
		return
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	fmt.Println("Program is running.")

	for {
		fmt.Println("Running...")
		time.Sleep(1 * time.Second)
	}
}
