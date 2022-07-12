package platform

import (
	"ProQueries/structs/block"
	"math/rand"
)

//const N = 10000
//const l = 10

type Rander struct {
	//db      stateDB.StateDB
	//headers Headers
	Height int // 块高度
	N      int // 账户数目
}

func NewRander(Height int, N int) *Rander {
	r := &Rander{
		Height: Height,
		N:      N,
	}
	return r
}

func RandAcconutKey() *block.AccountKey {
	key := block.NewAccountKey(rand.Intn(20))
	return key
}

func GetAcconutKey(i int) *block.AccountKey {
	key := block.NewAccountKey(i)
	return key
}

func (r *Rander) SetHeight(n int) {
	r.Height = n
}

func (r *Rander) RandAcconut() *block.Account {
	key := RandAcconutKey()
	ac := block.NewAcconunt(*key, rand.Intn(20), *block.NewnullPointer())
	return ac
}

func (r *Rander) RandTx(time int) *block.Tx {
	//num := rand.Intn(l)
	num1 := rand.Intn(r.N)
	num2 := rand.Intn(r.N)
	for num1 == num2 {
		num2 = rand.Intn(r.N)
	}
	to := GetAcconutKey(num1)
	from := GetAcconutKey(num2)
	ti := block.NewTiemstamp(time)
	tx := block.NewTx(block.TxState(1), *to, *from, rand.Int(), ti).(*block.Tx)
	return tx
}

func (r *Rander) RandTxs(i int) []block.Tx {
	var txs []block.Tx
	//var txs2 []int
	for j := 0; j < i; j++ {
		//k := rand.Int()
		tx := r.RandTx(j)
		txs = append(txs, *tx)
	}
	return txs
}
