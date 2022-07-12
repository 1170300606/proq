package block

type BasicBlock struct {
	header BasicHeader `json:"header"`
	datas  BasicDatas  `json:"data"`
}
type BasicHeader struct {
	BlockHash     []byte `json:"block_hash"`         //区块的信息
	LastBlockHash []byte `json:"last_block_hash"`    // 上一个区块的信息
	TxsHash       []byte `json:"txs_hash"`           // transactions
	Root          []byte `json:"account_state_hash"` // 状态数据库根
	Height        int    `json:"block_height""`      // 块高度
	//Bloom         Bloom `json:"blooom_filter""`  // 布隆过滤器
}

type BasicDatas struct {
	Txs []Tx `json:"txs"` // transcations
	//hash []byte // temp value
}

func NewBasicHeader(
	lastBlockHash []byte,
	txsHash []byte,
	root []byte,
	height int) *BasicHeader {
	header := &BasicHeader{
		//BlockHash:     blockHash,
		LastBlockHash: lastBlockHash,
		TxsHash:       txsHash,
		Root:          root,
		Height:        height,
	}
	return header
}

func NewBasicDatas(txs []Tx,
	hash []byte) *BasicDatas {

	datas := &BasicDatas{
		Txs: txs,
		//hash: hash,
	}
	return datas
}

func (header *BasicHeader) Sethash(hash []byte) {
	header.BlockHash = hash
}
