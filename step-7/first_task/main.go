package main

import (
	"errors"
	"fmt"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	ch := make(chan int, 1)

	go func(n int) {
		if n < 0 {
			ch <- -1
		}

		switch n {
		case 0:
			ch <- 0
		case 1:
			ch <- 1
		}
		num1 := 0
		num2 := 1

		for i := 1; i < n; i++ {
			num1, num2 = num2, num1+num2
		}

		ch <- num2
	}(n)

	select {
	case result := <-ch:
		if result < 0 {
			return 0, errors.New("n must be non-negative")
		}
		return result, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {
	fmt.Println(TimeoutFibonacci(10, 1*time.Second))
}
