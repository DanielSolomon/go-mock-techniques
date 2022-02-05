package mpmonkey

import "time"

type User struct {
	Name      string
	Age       uint8
	CreatedAt time.Time
}

func TestableCreateUser(name string, age uint8) User {
	return User{
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
	}
}
