package main

import (
	"fmt"
	"math"
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
	for val := range ch1 {
		valSq := int(math.Pow(float64(val), 2))
		ch2 <- valSq
	}
	close(ch2)
}
