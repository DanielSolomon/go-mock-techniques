package mockgen

import (
	"errors"
	"testing"

	"mocking/interfaces/mock/mockgen/database/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddUserAndAddresses(t *testing.T) {
	// Setup.
	ctrl := gomock.NewController(t)
	dbClient := mocks.NewMockIDatabaseClient(ctrl)

	id, name, age := "1337", "foo", uint8(17)
	dbClient.EXPECT().CreateUser(id, name, age).Return(errors.New("user already exists"))

	// Trigger.
	err := CreateUserWithAddresses(dbClient, id, name, age)

	// Validate.
	assert.NotNil(t, err, "expected an error")
}
