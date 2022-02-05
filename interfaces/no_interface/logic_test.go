package ni

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserAndAddresses(t *testing.T) {
	// Setup.
	dbClient := DatabaseClient{""} // cannot use NewDatabaseClient cause it will connect to the db.

	// Trigger.
	id := "1337"
	CreateUserWithAddresses(dbClient, id, "foo", 13) // cannot use it directly, it will try to insert to db.

	// Validate.
	user, _ := dbClient.FetchUser(id)
	assert.Equal(t, id, user.Id, "expected user")
}
