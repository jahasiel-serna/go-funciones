package main

import "fmt"

func sum(args ...int) int {
	bigger := args[0]
	for _, v := range args {
		if v > bigger {
			bigger = v
		}
	}
	return bigger
}

func main() {
	numbs := []int{10, 20, 30, 40, 80, 90, 100, 50, 60, 70}
	fmt.Println(sum(numbs...))
}