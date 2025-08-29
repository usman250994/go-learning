package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	intArrayDeclaration()

	stringArrayDeclaration()

	stringSplit("Hello, Go!")

	readInput()
	readFileWithErrHandling("file.txt")

}

func intArrayDeclaration() {
	myIntArr := []int{1, 2, 3, 4, 5}

	for i, v := range myIntArr {
		fmt.Println("Index:", i, "Value:", v)
	}
}

func stringArrayDeclaration() {
	myStringArr := []string{"one", "two", "three", "four", "five"}

	for i, v := range myStringArr {
		fmt.Println("Index:", i, "Value:", v)
	}
}

func stringSplit(s string) {

	myString := s

	splittedString := strings.Split(myString, "")
	for i, v := range splittedString {
		fmt.Println("Index:", i, "Value:", v)
	}

	joinedString := strings.Join(splittedString, "-")
	fmt.Println("Joined String:", joinedString)

	joinedSubString := strings.Join(splittedString[1:4], "-")
	fmt.Println("Joined Sub String:", joinedSubString)
}

// readInput reads lines from standard input and prints lines that occur more than once along with their counts.
func readInput() {

	fmt.Println("Enter lines of text. Press Ctrl+Z then Enter (Windows) or Ctrl+D (Unix/Mac) to finish:")

	counter := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counter[input.Text()]++
	}

	for line, count := range counter {
		if count > 1 {
			fmt.Printf("Line: %s, Count: %d\n", line, count)
		}
	}
}

// read file and return number of lines with error handling
func readFileWithErrHandling(filePath string) {

	counter := make(map[string]int)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for line, count := range counter {
		if count > 1 {
			fmt.Printf("Line: %s, Count: %d\n", line, count)
		}
	}
}
