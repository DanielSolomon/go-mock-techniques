package embed

import "errors"

type Address struct {
	Street string
	City   string
}

type User struct {
	Id        string
	Name      string
	Age       uint8
	Addresses []Address
}

type DatabaseClient struct {
	connectionString string
}

func NewDatabaseClient(connectionString string) DatabaseClient {
	// Open connection ...

	return DatabaseClient{connectionString: connectionString}
}

func (dc *DatabaseClient) CreateUser(id string, name string, age uint8) error {
	if id == "" {
		return errors.New("Name cannot be an empty string")
	}
	// dc.insert...
	return nil
}

func (dc *DatabaseClient) AddAddressToUser(id string, street string, city string) error {
	// dc.insert...
	return nil
}

func (dc *DatabaseClient) FetchUser(id string) (User, error) {
	// dc.select.Users.join.Addresses...
	return User{}, nil
}

func (dc *DatabaseClient) FetchUsersByAge(age uint8) ([]User, error) {
	return nil, nil
}

func (dc *DatabaseClient) FetchAddressesByCity(city string) ([]Address, error) {
	return nil, nil
}
