package block

import (
	"ProQueries/structs/merkletree"
	"ProQueries/structs/merkletree/datas"
)

type Txspool struct {
	database merkletree.BPlusTree
}

func (db *Txspool) Signall() {
	db.database.Signall()
}

func NewTxspool() *Txspool {
	var tree merkletree.BPlusTree
	(&tree).Initialize()
	st := &Txspool{
		database: tree,
	}
	return st
}

//
func (db *Txspool) Gethash() []byte {
	return db.database.GetRootHash()
}

func (db *Txspool) Insert(Key datas.Data_R, data *datas.Data_All) (*merkletree.BPlusFullNode, bool) {
	return db.database.Insert(Key, data)
}

func (db *Txspool) Delete(Key datas.Data_R) (*merkletree.BPlusFullNode, bool) {
	return db.database.Remove(Key)
}

func (db *Txspool) Find(Key datas.Data_R) (*datas.Data_All, bool) {
	return db.database.FindData(Key)
}

func (db *Txspool) query(left datas.Data_R, right datas.Data_R) datas.Vo {
	return db.database.RangeQUery(left, right)
}

func (tree *Txspool) RangeQUery(left datas.Data_R, right datas.Data_R) datas.Vo {
	return tree.database.RangeQUery(left, right)
}
