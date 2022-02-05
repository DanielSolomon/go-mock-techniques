package di

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name = "I'm Yogev"
	age  = uint8(25)
)

func TestUntestableCreateUser(t *testing.T) {
	// Trigger.
	name, age := "I'm Yogev", uint8(25)
	user := UntestableCreateUser(name, age)

	// Validate.
	assert.Equal(t, User{name, age, time.Now()}, user, "wrong created user")
}

func TestUntestableCreateUserAroundTime(t *testing.T) {
	// Trigger.
	user := UntestableCreateUser(name, age)

	// Validate.
	assert.Equal(t, name, user.Name, "wrong created user name")
	assert.Equal(t, age, user.Age, "wrong created user age")
	assert.InEpsilon(t, time.Now().UnixNano(), user.CreatedAt.UnixNano(), 0.1*float64(time.Second), "wrong created time")
}

func TestTestableCreateUser(t *testing.T) {
	// Setup.
	now := time.Now()
	nowFunc := func() time.Time {
		return now
	}

	// Trigger.
	user := TestableCreateUser(name, age, nowFunc)

	// Validate.
	assert.Equal(t, User{name, age, now}, user, "wrong created user")
}
