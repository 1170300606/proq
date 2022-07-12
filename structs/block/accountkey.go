package block

type AccountKey struct {
	Address int `json:"account_address"` //账户
}

func NewAccountKey(address int) *AccountKey {
	ac := &AccountKey{
		Address: address,
	}
	return ac
}

func RandAccountKey(address int) *AccountKey {
	ac := &AccountKey{
		Address: address,
	}
	return ac
}

func (ac *AccountKey) Islower(data_2 AccountKey) bool {

	return (ac.Address <= data_2.Address)
}

func (ac *AccountKey) Ishigher(data_2 AccountKey) bool {

	return (ac.Address >= data_2.Address)
}

func (ac *AccountKey) Equals(data_2 AccountKey) bool {

	return (ac.Address == data_2.Address)
}

func (ac *AccountKey) Isnull() bool {

	return false
}
