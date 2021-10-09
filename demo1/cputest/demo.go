package cputest

import (
	"fmt"
	"time"
)

// cpu intensive function
func delay() {
	startT := time.Now()

	total := 0
	for i:=1; i <= 3000000000; i++ {
		total += 1
	}

	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}
