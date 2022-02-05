package embed

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDatabaseClient struct {
	IDatabaseClient
	users map[string]User
}

func newMockDatabaseClient() mockDatabaseClient {
	return mockDatabaseClient{users: make(map[string]User)}
}

func (dc *mockDatabaseClient) CreateUser(id string, name string, age uint8) error {
	if id == "" {
		return errors.New("Name cannot be an empty string")
	}
	dc.users[id] = User{id, name, age, make([]Address, 0)}
	return nil
}

func (dc *mockDatabaseClient) AddAddressToUser(id string, street string, city string) error {
	user := dc.users[id]
	user.Addresses = append(user.Addresses, Address{street, city})
	return nil
}

func (dc *mockDatabaseClient) FetchUser(id string) (User, error) {
	return dc.users[id], nil
}

func TestAddUserAndAddresses(t *testing.T) {
	// Setup.
	dbClient := newMockDatabaseClient()

	// Trigger.
	id := "1337"
	CreateUserWithAddresses(&dbClient, id, "foo", 13) // cannot use it directly, it will try to insert to db.

	// Validate.
	user, _ := dbClient.FetchUser(id)
	assert.Equal(t, id, user.Id, "expected user")
}
