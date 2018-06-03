package models

import "fmt"

type User struct {
	ID       int
	UserName string
}

func SampleUsers() []User {
	users := make([]User, 0, 10)
	for i := 0; i < 10; i++ {
		users = append(users, User{ID: i, UserName: fmt.Sprint("testuser", i)})
	}
	return users
}
