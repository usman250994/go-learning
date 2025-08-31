package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gostory/user"
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

	userResult := readUsers()
	createNewUser(userResult)
}

func readUsers() user.User {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var users []user.User
	if err := json.Unmarshal(body, &users); err != nil {
		panic(err)
	}

	fmt.Printf("Number of users: %d\n", len(users))
	for _, user := range users {
		fmt.Println(user.GetInfo())
	}

	return users[0]
}

func createNewUser(u user.User) {
	newUser := user.NewUser(250994, u.Name+"-usman", u.Age+1, user.Address{})

	fmt.Printf("Creating new user: %v\n", newUser)

	jsonData, _ := json.Marshal(newUser)

	resp, err := http.Post("https://jsonplaceholder.typicode.com/users",
		"application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		panic(fmt.Sprintf("Failed to create user: %s", resp.Status))
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("User created successfully", string(body))
}
