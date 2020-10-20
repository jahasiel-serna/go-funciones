package main

import "fmt"

func oddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		var odd = i
		i += 2
		return odd
	}
}

func main() {
	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}