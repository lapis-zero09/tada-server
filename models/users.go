package models

import (
	"fmt"
)

type User struct {
	ID       int    `gorm:"primary_key" form:"id" json:"id"`
	UserName string `gorm:"not null" form:"userName" json:"userName"`
}

func SampleUsers() []User {
	users := make([]User, 0, 10)
	for i := 0; i < 10; i++ {
		users = append(users, User{UserName: fmt.Sprint("testuser", i)})
	}
	return users
}
