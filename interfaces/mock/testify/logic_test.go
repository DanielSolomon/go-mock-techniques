package mocktestify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDatabaseClient struct {
	mock.Mock
	IDatabaseClient
}

func (dc *mockDatabaseClient) CreateUser(id string, name string, age uint8) error {
	args := dc.Called(id, name, age)
	return args.Error(0)
}

func TestAddUserAndAddresses(t *testing.T) {
	// Setup.
	dbClient := new(mockDatabaseClient)

	id, name, age := "1337", "foo", uint8(17)
	dbClient.On("CreateUser", id, name, age).Return(errors.New("user already exists"))

	// Trigger.
	err := CreateUserWithAddresses(dbClient, id, name, age)

	// Validate.
	assert.NotNil(t, err, "expected an error")
}
