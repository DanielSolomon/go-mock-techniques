package embed

type IDatabaseClient interface {
	CreateUser(id string, name string, age uint8) error
	AddAddressToUser(id string, street string, city string) error
	FetchUser(id string) (User, error)
	FetchUsersByAge(age uint8) ([]User, error)
	FetchAddressesByCity(city string) ([]Address, error)
}

func CreateUserWithAddresses(dbClient IDatabaseClient, id string, name string, age uint8, addresses ...Address) error {
	err := dbClient.CreateUser(id, name, age)
	if err != nil {
		return err
	}
	for _, address := range addresses {
		err = dbClient.AddAddressToUser(id, address.Street, address.City)
		if err != nil {
			return err
		}
	}
	return nil
}
