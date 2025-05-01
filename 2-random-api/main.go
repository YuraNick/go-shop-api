package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Устанавливаем seed (например, текущее время)
	seed := time.Now().UnixNano()
	// Создаем новый источник случайных чисел
	source := rand.NewSource(seed)
	// Создаем новый генератор случайных чисел
	random := rand.New(source)

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		w.Write([]byte(strconv.Itoa(random.Intn(6) + 1)))
		return
	})

	port := 81
	fmt.Printf("server is running http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
