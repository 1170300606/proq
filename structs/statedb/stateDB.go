package statedb

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree"
	"ProQueries/structs/merkletree/datas"
)

type Position *merkletree.BPlusFullNode
type StateDB struct {
	database merkletree.BPlusTree
}

func NewstateDB() *StateDB {
	var tree merkletree.BPlusTree
	(&tree).Initialize()
	st := &StateDB{
		database: tree,
	}
	return st
}

func (db *StateDB) Gethash() []byte {
	return db.database.GetRootHash()
}

func (db *StateDB) Insert(Key datas.Data_R, data *datas.Data_All) (*merkletree.BPlusFullNode, bool) {
	return db.database.Insert(Key, data)
}

func (db *StateDB) Delete(Key datas.Data_R) (*merkletree.BPlusFullNode, bool) {
	return db.database.Remove(Key)
}

func (db *StateDB) Find(Key datas.Data_R) (*datas.Data_All, bool) {
	return db.database.FindData(Key)
}

func (db *StateDB) query(left datas.Data_R, right datas.Data_R) datas.Vo {
	return db.database.RangeQUery(left, right)
}

func (db *StateDB) change(Key datas.Data_R, data *datas.Data_All) (*merkletree.BPlusFullNode, bool) {
	//return db.database.Insert(Key, data)
	//db.Delete(Key)
	return db.Insert(Key, data)
}

func (db *StateDB) Signall() {
	db.database.Signall()
}

func (db *StateDB) Change(state *datas.Data_All, value int) bool {
	ac := state.Showdata().Showdata().(*block.Account)
	ac.Change(value)
	return true
}

func (db *StateDB) Transfer(thetx block.Tx, order int, height int) block.Tx {
	//TODO 转账

	to := datas.NewDataR(block.NewAcconunt(thetx.To, 0, block.Pointers{}))
	from := datas.NewDataR(block.NewAcconunt(thetx.From, 0, block.Pointers{}))
	//var to_ac *block.Account
	value := thetx.Value
	hash := thetx.Hash()
	pointer := block.NewPointer(height, order, hash)
	toState, err := db.Find(*to)
	if !err {

	}

	account := toState.Showdata().Showdata().(*block.Account)
	po_1 := account.Getpointer()
	account.SetPointer(*pointer)
	db.Change(toState, value)

	fromState, err := db.Find(*from)
	if !err {

	}
	account = fromState.Showdata().Showdata().(*block.Account)
	po_2 := account.Getpointer()
	account.SetPointer(*pointer)
	db.Change(fromState, -value)

	//to.Change(value)
	//from.Change(-value)
	//to_r := datas.NewDataR(&to)
	//from_r := datas.NewDataR(&from)
	//to_all := datas.NewDataAll(to_r)
	//from_all := datas.NewDataAll(from_r)
	//db.change(*to_r, to_all)
	//db.change(*from_r, from_all)
	//var answer tx.Tx
	Po := block.NewnullPointer()
	Po.Addpo(po_1)
	Po.Addpo(po_2)
	thetx.Addpointer(*Po)
	return thetx
}

func (db *StateDB) RangeQUery(left datas.Data_R, right datas.Data_R) datas.Vo {
	return db.database.RangeQUery(left, right)
}
