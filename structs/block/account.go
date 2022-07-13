package block

import (
	"ProQueries/structs/merkletree/datas"
	"github.com/tendermint/tendermint/libs/json"
)

type Account struct {
	AcKey   AccountKey `json:"account_key"`   //账户
	Value   int        `json:"account_value"` //余额
	Pointer Pointer    `json:"protx_pointer"` //指向与账户关联的的上一条交易的交易指针
}

func NewAcconunt(account AccountKey, value int, pointer Pointers) *Account {
	ac := &Account{
		AcKey:   account,
		Value:   value,
		Pointer: Pointer{},
	}
	return ac
}

func (ac *Account) Islower(data_2 datas.Content) bool {
	data2 := data_2.(*Account)
	return ac.AcKey.Islower(data2.AcKey)
}

func (ac *Account) Is_lower(data_2 AccountKey) bool {

	return ac.AcKey.Islower(data_2)
}

func (ac *Account) Change(v int) {
	ac.Value = ac.Value + v
}

func (ac *Account) Ishigher(data_2 datas.Content) bool {
	data2 := data_2.(*Account)
	return ac.AcKey.Ishigher(data2.AcKey)
}

func (ac *Account) Is_higher(data_2 AccountKey) bool {
	//data2 := data_2.(*Account)
	return ac.AcKey.Ishigher(data_2)
}

func (ac *Account) Equals(data_2 datas.Content) bool {
	data2 := data_2.(*Account)
	return ac.AcKey.Equals(data2.AcKey)
}

func (ac *Account) Equal(data_2 AccountKey) bool {
	//data2 := data_2.(*Account)
	return ac.AcKey.Equals(data_2)
}

func NewNullAcconunt() *Account {
	ac := &Account{}
	return ac
}

func (ac *Account) SetPointer(po Pointer) bool {
	ac.Pointer = po
	//TODO
	return true
}

func (ac *Account) Isnull() bool {

	return false
}

func (ac *Account) Tobyte() []byte {
	a, _ := json.Marshal(ac)
	return a
}

func (ac *Account) Getpointer() Pointer {

	return ac.Pointer
}
