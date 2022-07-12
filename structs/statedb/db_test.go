package statedb

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree/datas"

	"testing"
)

func TestDb(t *testing.T) {
	//r := NewRander(0, 100)
	db := NewstateDB()
	//txs := r.RandTxs(100)
	for i := 0; i <= 1000000; i++ {
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		r := datas.NewDataR(account)
		all := datas.NewDataAll(r)
		db.Insert(*r, all)
	}
}
