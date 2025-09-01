package main

import "fmt"

func main() {

	res := add(1, 2)
	res2 := multiply(2, 3)
	fmt.Printf("res 1: %d res 2: %d\n", res, res2)

}

// simple add to use in tests
func add(a int, b int) int {
	return a + b
}

// simple multiply to use in tests
func multiply(a int, b int) int {
	result := 0
	for i := 0; i < b; i++ {
		result += a
	}
	return result
}
