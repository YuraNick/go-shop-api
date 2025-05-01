package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go generate(ch1)
	go readAndSquare(ch1, ch2)

	for val := range ch2 {
		fmt.Print(val, " ")
	}
}

func generate(ch chan int) {
	const len int = 10
	const maxVal int = 100
	slice := make([]int, len)
	for i := range slice {
		slice[i] = rand.Intn(maxVal)
		ch <- slice[i]
	}
	close(ch)
}

func readAndSquare(ch1, ch2 chan int) {
	i := 0
	for val := range ch1 {
		valSq := val * val
		ch2 <- valSq
		i++
		if i == 10 {
			break
		}
	}
	close(ch2)
}
