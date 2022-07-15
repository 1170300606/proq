package platform

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree/datas"
	"ProQueries/structs/statedb"
	"fmt"
	"testing"
	time2 "time"
)

func TestNewTxs(t *testing.T) {
	node := NewNode()
	tx := node.NewTxs()

	if len(tx) == 0 {
		t.Errorf("wrong")
	}
}

func TestNewBlock(t *testing.T) {
	node := NewNode()
	node.NewState()
	block := node.NewBlock()
	block.Show()
}

func TestFind(t *testing.T) {
	node := NewNode()
	node.NewState()
	block_ := node.NewBlock()
	txs := block_.Datas.Txs
	_, err, _ := node.FindState(*block.NewAcconunt(txs[0].To, 0, block.Pointers{}))
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDb(t *testing.T) {
	//r := NewRander(0, 100)
	db := statedb.NewstateDB()
	//txs := r.RandTxs(100)
	for i := 0; i <= 10000; i++ {
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		r := datas.NewDataR(account)
		all := datas.NewDataAll(r)
		db.Insert(*r, *all)
	}
}

func TestPool(t *testing.T) {
	r := NewRander(0, 100)
	txs := r.RandTxs(100)
	txpool := statedb.NewTxspool()
	for i := 0; i < len(txs); i++ {
		//fmt.Println(i)
		r, all := makeTxPoint(txs[i])
		txpool.Insert(*r, *all)
	}
}

func TestNewnodes(t *testing.T) {
	node := NewNode()
	node.NewBlocks(10)
	if node.Height != 9 {
		t.Errorf("wrong")
	}
	//var a []int
	//a = append(a, 1)
	//a = append(a, 2)
	//a = append(a, 3)
}

func TestBackward(t *testing.T) {
	t0 := time2.Now()
	node := NewNode()
	node.NewBlocks(100)
	elapsed := time2.Since(t0)
	fmt.Println("建立数据的时间是：", elapsed)
	key := *block.NewAccountKey(100)
	txs, vo := node.Backward(key)
	t1 := time2.Now()

	err := vo.Verifiable()
	elapsed = time2.Since(t1)
	fmt.Println("验证Vo的时间是：", elapsed)
	if !err {
		t.Errorf("wrong")
	}
	fmt.Println(txs)
	elapsed = time2.Since(t0)
	fmt.Println("总时间是：", elapsed)
}

func TestTime(t *testing.T) {
	t1 := time2.Now()

	//err := vo.Verifiable()
	//t2 := time2.Now()
	//fmt.Println("s")
	s := time2.Since(t1)
	fmt.Println(s)
}

func Test(t *testing.T) {
	fmt.Println(int(^uint(0) >> 1))
}
