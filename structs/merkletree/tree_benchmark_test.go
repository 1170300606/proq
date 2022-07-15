package merkletree

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree/datas"
	"fmt"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	t := NewTree()

	// insert b.N times
	for n := 0; n < b.N; n++ {
		account := block.NewAcconunt(*block.NewAccountKey(n), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)

		err := t.Insert(*key, *value)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}
}

func BenchmarkInsertFind(b *testing.B) {
	t := NewTree()

	// insert b.N times
	//for n := 0; n < b.N; n++ {
	for n := 0; n < 100; n++ {
		account := block.NewAcconunt(*block.NewAccountKey(n), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)

		err := t.Insert(*key, *value)

		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	// find one by one
	//for n := 0; n < b.N; n++ {
	for n := 0; n < 100; n++ {
		account := block.NewAcconunt(*block.NewAccountKey(n), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		key := datas.NewDataR(account)
		//value := datas.NewDataAll(key)
		_, err, _ = t.Find(*key, false)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}
}

func BenchmarkInsertDelete(b *testing.B) {
	t := NewTree()

	// insert b.N times
	for n := 0; n < b.N; n++ {
		account := block.NewAcconunt(*block.NewAccountKey(n), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)

		err := t.Insert(*key, *value)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}

	// delete them
	for n := 0; n < b.N; n++ {
		account := block.NewAcconunt(*block.NewAccountKey(n), 0, block.Pointers{})
		//acc := account.(datas.Data_R)
		key := datas.NewDataR(account)
		//value := datas.NewDataAll(key)
		err = t.Delete(*key)
		if err != nil {
			fmt.Printf("error: %s\n\n", err)
		}
	}
}
