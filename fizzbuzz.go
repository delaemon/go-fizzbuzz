package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 30; i++ {
		m := make([]byte, 0, 100)
		m = append(m, strconv.Itoa(i)...)
		if i%3 == 0 {
			m = append(m, " fizz"...)
		}
		if i%5 == 0 {
			m = append(m, " buzz"...)
		}
		fmt.Println(string(m))
	}
}
