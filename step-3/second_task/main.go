package main

func Receive(ch chan int) int {
	result := <- ch
	return result
}