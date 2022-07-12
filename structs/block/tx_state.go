package block

type TxState uint8

const (
	// tx status
	make  = TxState(0) // 创建账户
	tx    = TxState(1) // 正常交易
	delet = TxState(2) //注销账户
)
