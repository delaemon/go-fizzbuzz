package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 30; i++ {
		m := make([]byte, 0, 100)
		m = append(m, strconv.Itoa(i)...)
		if i%3 == 0 {
			m = append(m, " fizz"...)
		}
		if i%5 == 0 {
			m = append(m, " buzz"...)
		}
		c := i % 7 + 1
		fmt.Printf("\033[3%dm%s\n", c, string(m))
	}
}
