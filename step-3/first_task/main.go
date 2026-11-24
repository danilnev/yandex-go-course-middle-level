package main

func send(ch chan int, num int) {
	ch <- num
}