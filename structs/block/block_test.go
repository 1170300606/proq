package block

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	"math/rand"
	"testing"
)

func GetAcconutKey(i int) *AccountKey {
	key := NewAccountKey(i)
	return key
}

func TestTx(t *testing.T) {
	to := GetAcconutKey(1)
	from := GetAcconutKey(2)
	ti1 := NewTiemstamp(0)
	ti2 := NewTiemstamp(1)
	tx := NewTx(TxState(1), *to, *from, rand.Int(), ti1).(*Tx)
	tx2 := NewTx(TxState(1), *to, *from, rand.Int(), ti2).(*Tx)
	s, _ := json.Marshal(tx)
	s2, _ := json.Marshal(tx2)
	str := string(s)
	str2 := string(s2)
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str == str2)
}

func Test(t *testing.T) {
	to := GetAcconutKey(1)
	from := GetAcconutKey(2)
	ti1 := NewTiemstamp(0)
	ti2 := NewTiemstamp(1)
	tx := NewTx(TxState(1), *to, *from, rand.Int(), ti1).(*Tx)
	tx2 := NewTx(TxState(1), *to, *from, rand.Int(), ti2).(*Tx)
	s, _ := json.Marshal(tx)
	s2, _ := json.Marshal(tx2)
	str := string(s)
	str2 := string(s2)
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str == str2)
}
