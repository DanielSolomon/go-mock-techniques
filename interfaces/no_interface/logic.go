package ni

func CreateUserWithAddresses(dbClient DatabaseClient, id string, name string, age uint8, addresses ...Address) error {
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
