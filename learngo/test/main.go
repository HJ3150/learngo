package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	m, _ := time.ParseDuration("1m30s")
	for i := 0; i < 1000; i++ {
		fmt.Printf("Take off in t-%.0f seconds.\n", m.Seconds())
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Time:", elapsedTime)
}
