package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Устанавливаем seed (например, текущее время)
	seed := time.Now().UnixNano()
	// Создаем новый источник случайных чисел
	source := rand.NewSource(seed)
	// Создаем новый генератор случайных чисел
	random := rand.New(source)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go generate(ch1, random)
	go readAndSquare(ch1, ch2)

	for val := range ch2 {
		fmt.Print(val, " ")
	}
}

func generate(ch chan int, rnd *rand.Rand) {
	const len int = 10
	const maxVal int = 100
	slice := make([]int, len)
	for i := range slice {
		slice[i] = rnd.Intn(maxVal)
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
