package user

import "fmt"

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}
type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func NewUser(id int, name string, age int, address Address) *User {
	return &User{
		ID:      id,
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (u User) GetInfo() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Address: %+v", u.ID, u.Name, u.Age, u.Address)
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
