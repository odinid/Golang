package main

import "fmt"

func fibonacci() func() int {
	var x1, x2, sum = -1, 1, 0

	return func() int {
		sum = x1 + x2
		x1 = x2
		x2 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
