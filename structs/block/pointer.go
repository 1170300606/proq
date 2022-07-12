package block

import "reflect"

type Pointer struct { // 注意索引大小
	Height   int    `json:"block_height"` // 交易所在块高度
	Position int    `json:"tx_position"`  // 是偏移量 int
	Hash     []byte `json:"tx_hash"`      // 交易dig
}

type Pointers struct { // 注意索引大小
	Pointers []Pointer `json:"Pointers"` // 与之相关的一组交易
}

func NewnullPointer() *Pointers {
	po := &Pointers{}
	return po
}

func (Pos *Pointers) Addpo(po Pointer) {
	Pos.Pointers = append(Pos.Pointers, po)
}

func NewPointer(height int, position int, hash []byte) *Pointer {
	po := &Pointer{
		Height:   height,
		Position: position,
		Hash:     hash,
	}
	return po
}

func (a Pointer) IsEmpty() bool {
	return reflect.DeepEqual(a, Pointer{})
}

func (a Pointers) IsEmpty() bool {
	return reflect.DeepEqual(a, Pointers{})
}
