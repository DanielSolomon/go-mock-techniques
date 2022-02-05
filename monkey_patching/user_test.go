package mp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name = "I'm Yogev"
	age  = uint8(25)
	now  = time.Now()
)

func patchNow() {
	nowFunc = func() time.Time {
		return now
	}
}

func TestTestableCreateUser(t *testing.T) {
	// Setup.
	patchNow()

	// Trigger.
	user := TestableCreateUser(name, age)

	// Validate.
	assert.Equal(t, User{name, age, now}, user, "wrong created user")
}

func TestAfter(t *testing.T) {
	fmt.Println(nowFunc())
	time.Sleep(time.Second)
	fmt.Println(nowFunc())
}
