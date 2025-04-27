package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		w.Write([]byte(strconv.Itoa(rand.Intn(6) + 1)))
		return
	})
	port := 81
	fmt.Printf("server is running http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
