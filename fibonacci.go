package main

import "fmt"

func fibonacci(pos int64) int64 {
	if pos == 1 || pos == 0 {
		return pos
	} else {
		return fibonacci(pos - 1) + fibonacci(pos - 2)
	}
}

func main() {
	var pos int64
	fmt.Scan(&pos)
	fmt.Println(fibonacci(pos))
}