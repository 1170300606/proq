package statedb

import (
	"ProQueries/structs/merkletree"
	"ProQueries/structs/merkletree/datas"
)

type Txspool struct {
	//database merkletree.BPlusTree
	database merkletree.Tree
}

func (db *Txspool) Signall() {
	db.database.SignAll()
}

func NewTxspool() *Txspool {
	var tree merkletree.Tree
	//(&tree).Initialize()
	st := &Txspool{
		database: tree,
	}
	return st
}

//
func (db *Txspool) Gethash() []byte {
	return db.database.GetRootHash()
}

func (db *Txspool) Insert(Key datas.Data_R, data datas.Data_All) error {
	return db.database.Insert(Key, data)
}

func (db *Txspool) Delete(Key datas.Data_R) error {
	return db.database.Delete(Key)
}

func (db *Txspool) Find(Key datas.Data_R) (*datas.Data_All, error) {
	re, err, _ := db.database.Find(Key, false)
	return &re.Value, err
}

//func (db *Txspool) query(left datas.Data_R, right datas.Data_R) datas.Vo {
//	return db.database.RangeQUery(left, right)
//}
//
//func (tree *Txspool) RangeQUery(left datas.Data_R, right datas.Data_R) datas.Vo {
//	return tree.database.RangeQUery(left, right)
//}
