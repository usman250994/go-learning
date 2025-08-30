package user

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func NewUser(id int, name string, age int) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}

func (u User) GetInfo() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d", u.ID, u.Name, u.Age)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}
