package mp

import "time"

type User struct {
	Name      string
	Age       uint8
	CreatedAt time.Time
}

func UntestableCreateUser(name string, age uint8) User {
	return User{
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
	}
}

var nowFunc = time.Now

func TestableCreateUser(name string, age uint8) User {
	return User{
		Name:      name,
		Age:       age,
		CreatedAt: nowFunc(),
	}
}
