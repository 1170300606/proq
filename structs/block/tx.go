package block

import (
	"ProQueries/structs/merkletree/datas"
	"crypto/sha256"
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
)

type Txs []Tx
type Tx struct {
	txType  TxState    `json:"tx_txtype"`  // 显示交易的类型
	To      AccountKey `json:"tx_to"`      // 收款账户
	From    AccountKey `json:"tx_from"`    // 付款账户
	Value   int        `json:"tx_value"`   // 交易额
	Pointer Pointers   `json:"tx_pointer"` // 指向交易链条中的上一条交易们
	Time    Tiemstamp  `json:"tx_time"`    // 交易产生的时间戳
}

func (tx *Tx) Tobyte() []byte {
	a, _ := json.Marshal(tx)
	return a
}

func NewTx(txType TxState, to AccountKey, from AccountKey, value int, time Tiemstamp) datas.Content {
	tx := &Tx{
		txType: txType,
		To:     to,
		From:   from,
		Value:  value,
		Time:   time,
		//pointer: pointer,
	}
	return tx
}

func NewNullTx() *Tx {
	tx := &Tx{}
	return tx
}

func (tx *Tx) Addpointer(pointer Pointers) {
	tx.Pointer = pointer
}

func (tx *Tx) Islower(data_2 datas.Content) bool {
	tx2 := data_2.(*Tx)
	//s, _ := json.Marshal(tx)
	//s2, _ := json.Marshal(tx2)
	//str := string(s)
	//str2 := string(s2)
	////return tx.To.Islower(tx2.To)
	//return (str <= str2)
	return tx.Time.Islower(tx2.Time)
	//return false
}

func (tx *Tx) Ishigher(data_2 datas.Content) bool {
	tx2 := data_2.(*Tx)
	//s, _ := json.Marshal(tx)
	//s2, _ := json.Marshal(tx2)
	//str := string(s)
	//str2 := string(s2)
	////return tx.To.Islower(tx2.To)
	//return (str >= str2)
	return tx.Time.Ishigher(tx2.Time)
	//return tx.To.Ishigher(tx2.To)
	//return false
}

func (tx *Tx) Hash() []byte {
	h := sha256.New()
	s, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	//str := string(s)
	//io.WriteString(h, s)
	h.Write(s)
	a := h.Sum(nil)
	return a
}

func (tx *Tx) Equals(data_2 datas.Content) bool {
	//TODO 需要修改
	tx2 := data_2.(*Tx)
	//x2 := data_2.(*Tx)
	//s, _ := json.Marshal(tx)
	//s2, _ := json.Marshal(tx2)
	//str := string(s)
	//str2 := string(s2)
	////return tx.To.Islower(tx2.To)
	//return (str == str2)
	return tx.Time.Equals(tx2.Time)
	//	return false
}

func (tx *Tx) Isnull() bool {

	return false
}
