package platform

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree"
	"ProQueries/structs/merkletree/datas"
	"ProQueries/structs/statedb"
	"crypto/sha256"
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	time2 "time"
)

const N = 10000
const n = 100

type Headers []block.BasicHeader

type Node struct {
	db *statedb.StateDB
	//headers    Headers
	Blocks     []block.ProBlock
	r          *Rander
	Height     int // 当前块高度
	lastBkHash []byte
	Pool       *block.Txspool
}

/*
	生成一个新的节点
*/

func NewNode() *Node {
	db := statedb.NewstateDB()
	var blocks []block.ProBlock
	r := NewRander(0, N)
	var lastBkHash []byte
	pool := block.NewTxspool()
	node := &Node{
		db:         db,
		Blocks:     blocks,
		r:          r,
		lastBkHash: lastBkHash,
		Pool:       pool,
		Height:     -1,
	}
	node.NewStateby(N)
	return node
}

/*
	随机生成区块的交易部分，得到root
	执行生成的交易，更新状态数据库
	生成区块头部分
*/
func (node *Node) NewBlocks(n int) {
	var time time2.Duration = 0
	for i := 0; i < n; i++ {
		problock := node.NewBlock()
		t0 := time2.Now()
		node.db.Signall()
		elapsed := time2.Since(t0)
		time = time + elapsed
		node.Blocks = append(node.Blocks, *problock)
		//node.Height++
	}

	fmt.Println("更新数据的时间是：", time)
}

func (node *Node) addheight() {
	node.Height++
	node.r.SetHeight(node.Height)
}

func (node *Node) NewBlock() *block.ProBlock {
	node.addheight()           //Heigjt加一
	txs := node.NewTxs()       //随机生成一组交易
	txs = node.ExecuteTxs(txs) //执行交易
	txs = node.NewDatas(txs)   //将交易加入交易pool中
	datas := block.NewBasicDatas(txs, nil)
	header := *node.NewHeader(node.Pool.Gethash())
	newblock := block.NewProBlock(header, *datas)
	node.Sethash(newblock)
	//node.Blocks = append(node.Blocks, *newblock)
	return newblock
}

func (node *Node) NewPool() {
	node.Pool = block.NewTxspool()
}

/*
	随机生成一组开户交易，用来初始化账户状态
*/
func (node *Node) NewState() {
	for i := 0; i <= 20; i++ {
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		r := datas.NewDataR(account)
		all := datas.NewDataAll(r)
		node.db.Insert(*r, all)
	}
}

/*
	随机生成一组开户交易，用来初始化账户状态
*/
func (node *Node) NewStateby(num int) {
	for i := 0; i <= num; i++ {
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		r := datas.NewDataR(account)
		all := datas.NewDataAll(r)
		node.db.Insert(*r, all)
	}
}

/*
	为一个新生的区块添加hash
*/
func (node *Node) Sethash(block *block.ProBlock) {
	h := sha256.New()
	s, err := json.Marshal(*block)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	//str := string(s)
	//io.WriteString(h, s)
	h.Write(s)
	a := h.Sum(nil)
	block.Sethash(a)
}

/*
	接受之前得到一组交易，将交易构成默克尔树，得到root
*/
func (node *Node) NewDatas(txs block.Txs) block.Txs {
	//txs := node.r.RandTxs(10)
	//root := []byte("")
	//TODO 将txs塞进M树中，得到root
	node.NewPool()
	for i := 0; i < len(txs); i++ {
		//fmt.Println(i)
		//if i == 34 {
		//	i = 34
		//}
		r, all := makeTxPoint(txs[i])
		node.Pool.Insert(*r, all)
	}
	node.Pool.Signall()
	//datas := block.NewBasicDatas(txs, nil)
	return txs
}

/*
	生成一组随机的交易
*/

func (node *Node) NewTxs() block.Txs {
	txs := node.r.RandTxs(n)
	return txs
}

/*
	执行交易后得到最新版本的DB
*/
func (node *Node) ExecuteTxs(txs block.Txs) block.Txs {
	//txs := datas.Txs
	var answer block.Txs
	for i := 0; i < len(txs); i++ {
		if i == 7 {
			i = 7
		}
		tx := node.db.Transfer(txs[i], i, node.Height)
		answer = append(answer, tx)
	}
	node.db.Signall()
	return answer
}

/*
	输入：交易集合的默克尔树根
	在执行交易后得到最新版本的DB，将DB的root
*/

func (node *Node) NewHeader(hash []byte) *block.BasicHeader {
	//TODO 需要更新后在生成新的header
	header := block.NewBasicHeader(node.lastBkHash, hash, node.db.Gethash(), node.Height)
	return header
}

//由于是模拟，只有一个节点，该节点既是查询发出的轻节点又是负责查询的全节点
func (node *Node) Backward(key block.AccountKey) (block.Txs, datas.Vo) {
	//TODO
	txs, vo := node.linQuery(key)
	return txs, vo
}

func (node *Node) Insert(Key datas.Data_R, data *datas.Data_All) (*merkletree.BPlusFullNode, bool) {
	return node.db.Insert(Key, data)
}

func (node *Node) Delete(Key datas.Data_R) (*merkletree.BPlusFullNode, bool) {
	return node.db.Delete(Key)
}

func (node *Node) FindState(Key block.Account) (*datas.Data_All, bool) {
	key := datas.NewDataR(&Key)
	return node.db.Find(*key)
}

//验证用的vo应该是交易默克尔树的Vo,而不是状态数据库默克尔树的Vo
func (node *Node) linQuery(key block.AccountKey) (block.Txs, datas.Vo) {
	//TODO
	var txs block.Txs
	tempAccount := datas.NewDataR(block.NewAcconunt(key, 0, block.Pointers{}))
	all, _ := node.db.Find(*tempAccount)
	Account := all.Showdata().Showdata().(*block.Account)
	po := Account.Pointer

	tx := node.gettx(po)
	//TODO Vo 没过去
	t1 := time2.Now()
	vo := node.db.RangeQUery(*tempAccount, *tempAccount)
	elapsed := time2.Since(t1)

	fmt.Println("建立Vo的时间是：", elapsed)
	t2 := time2.Now()
	//fmt.Println(ti2.Sub(ti1).Nanoseconds,)
	for tx != nil {
		txs = append(txs, *tx)
		if tx.Pointer.IsEmpty() {
			break
		}
		if tx.To.Equals(key) {
			po = tx.Pointer.Pointers[0]
			tx = node.gettx(po)
		} else {
			po = tx.Pointer.Pointers[1]
			tx = node.gettx(po)
		}
	}
	elapsed = time2.Since(t2)
	fmt.Println("查找txs的时间是：", elapsed)
	//vo := node.Pool.RangeQUery(*tempAccount,*tempAccount)

	return txs, vo
}

func (node *Node) gettx(po block.Pointer) *block.Tx {
	if po.IsEmpty() {
		return nil
	}
	tx := node.Blocks[po.Height].Datas.Txs[po.Position]
	return &tx
}

func (node *Node) ver(txs block.Txs, Vo datas.Vo) bool {
	//TODO
	return true
}

func makeTxPoint(tx block.Tx) (Key *datas.Data_R, data *datas.Data_All) {
	r := datas.NewDataR(&tx)
	all := datas.NewDataAll(r)
	return r, all
}
