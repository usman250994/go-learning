package main

import (
	"fmt"
	"os"
)

func main() {
	simpleForLoop(5)
	whileLoop(5)
	iterateThroughRange([]int{5, 4, 3, 2, 1})
	iterateThroughVariableRange(5)
	iterateThroughOsArgs()
}

func simpleForLoop(count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("Iteration:", i)
	}
}

func whileLoop(count int) {
	i := 1
	for i <= count {
		fmt.Println("Iteration:", i)
		i += 1
	}
}

// do not forget to give cmd line args to the execution
func iterateThroughRange(count []int) {
	for i, arg := range count {
		fmt.Println("Iteration:", i, arg)
		i += 1
	}
}

func iterateThroughVariableRange(count int) {
	values := []int{5, 4, 3, 2, 1}
	for i, arg := range values[:count] {
		fmt.Println("Iteration:", i, arg)
		i += 1
	}
}

func iterateThroughOsArgs() {
	for i, arg := range os.Args[1:] {
		fmt.Println("Iteration:", i, arg)
		i += 1
	}

}
