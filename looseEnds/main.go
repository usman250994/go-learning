package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := switchCase()
	fmt.Print(x)
}

// defining enum
type CoinFace string

// assigning enum values
const (
	Head    CoinFace = "head"
	Tail    CoinFace = "tail"
	Unknown CoinFace = "unknown"
)

// returning type saefe and clean values of coinface

func switchCase() CoinFace {
	// using math library for random number 0, 1
	x := rand.Intn(2) // or say +1 for 1 or 2
	switch x {
	case 1:
		return Head
	case 0:
		return Tail
		// hypothetical situation
	default:
		return Unknown
	}
}
