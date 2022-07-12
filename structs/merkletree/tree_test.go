package merkletree

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree/datas"

	//"ProQueries/structs/merkletree/datas"
	"ProQueries/structs/platform"
	"testing"
)

func makeTxPoint(tx block.Tx) (Key *datas.Data_R, data *datas.Data_All) {
	r := datas.NewDataR(&tx)
	all := datas.NewDataAll(r)
	return r, all
}

func Test(t *testing.T) {
	r := platform.NewRander(0, 100)
	txs := r.RandTxs(100)
	txpool := block.NewTxspool()
	for i := 0; i < len(txs); i++ {
		r, all := makeTxPoint(txs[i])
		txpool.Insert(*r, all)
	}
}
