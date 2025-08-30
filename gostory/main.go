package main

import (
	"fmt"
	"io"
	"net/http"
)

// find some simple way to use interface
func main() {

	//understand defer and recover together
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	//marshal into []user struct
	//unmarshal to check
	//call functions to update

}
