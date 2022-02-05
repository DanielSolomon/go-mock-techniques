package mpmonkey

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

var (
	name = "I'm Yogev"
	age  = uint8(25)
	now  = time.Now()
)

func monkeyPatchNow() {
	monkey.Patch(time.Now, func() time.Time { return now })
}

func TestPatchTestableCreateUser(t *testing.T) {
	// Setup.
	monkeyPatchNow()
	defer monkey.UnpatchAll()

	// Trigger.
	user := TestableCreateUser(name, age)

	// Validate.
	assert.Equal(t, User{name, age, now}, user, "wrong created user")
}
