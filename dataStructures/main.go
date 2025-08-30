package main

import (
	"fmt"
)

func main() {

	indexOfString("hello", 1)

	Constants()

	hashMap()

	autoSizeArr()
}

func indexOfString(word string, i int) {
	// prints byte at position i
	fmt.Printf("byte: %v\n", word[i])
	// prints actual character value
	fmt.Printf("char: %v\n", word[i:i+1])

}

func Constants() {
	type Celsius int
	type Fahrenheit float64

	const (
		celsius    Celsius    = 23
		fahrenheit Fahrenheit = 23.0
	)

	// Compiler error
	// invalid operation: Celsius == Fahrenheit
	// (mismatched types Celsius and Fahrenheit)
	// code below wont compile therefore commented

	//--> if Celsius == Fahrenheit

	// This conversion is safer
	if fahrenheit == Fahrenheit(celsius) {
		fmt.Printf("%v\n", celsius)
	}

	// use case of iota when declaring series of incrementing constants
	// iota is reset to 0 whenever the keyword const appears
	type weekday int

	const (
		Sunday    weekday = iota // 0
		Monday                   // 1
		Tuesday                  // 2
		Wednesday                // 3
		Thursday                 // 4
		Friday                   // 5
		Saturday
	)

	var today weekday = Friday
	fmt.Printf("Today is %vth day\n", today)

	//same numbers as previous constants example below

	type weekday_s int

	const (
		Sunday_s weekday_s = 1
		Monday_s
		Tuesday_s weekday_s = 3
		Wednesday_s
		Thursday_s weekday_s = 5
		Friday_s
		Saturday_s
	)

	fmt.Printf("same values copied  %v,%v, %v\n", Monday_s, Wednesday_s, Friday_s)
}

func hashMap() {

	var myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	// check if key exists
	val, ok := myMap["four"]
	if !ok {
		fmt.Printf("Key 'four' does not exist\n")
	} else {
		fmt.Printf("Key 'four' exists with value: %v\n", val)
	}

}

func autoSizeArr() {
	myArr := [...]string{1: "1", 2: "2", 5: "99"}

	for i := 0; i < len(myArr); i++ {
		fmt.Printf(" %v: , %v\n ", i, myArr[i])
	}
}
