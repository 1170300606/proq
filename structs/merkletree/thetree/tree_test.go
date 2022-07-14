package thetree

import (
	"ProQueries/structs/block"
	"ProQueries/structs/merkletree/datas"
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	"reflect"
	"testing"
)

func hello() {
	fmt.Println("bptree says 'hello friend'")
}

func TestInsertNilRoot(t *testing.T) {
	tree := NewTree()
	hello()

	//key := 1
	//value := []byte("test")
	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)

	err := tree.Insert(*key, *value)

	if err != nil {
		t.Errorf("%s", err)
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}
}

func TestInsertList(t *testing.T) {
	tree := NewTree()
	for i := 100; i > 0; i-- {
		//key := i
		//value := []byte("test")
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)
		fmt.Println(i)
		err := tree.Insert(*key, *value)
		if err != nil {
			t.Errorf("%s", err)
		}

		r, err, _ := tree.Find(*key, false)
		if err != nil {
			t.Errorf("%s\n", err)
		}

		if r == nil {
			t.Errorf("returned nil \n")
		}

		//if !reflect.DeepEqual(r.Value, value) {
		//if !reflect.DeepEqual(r.Value, value) {
		if !r.Value.Equals(value) {
			t.Errorf("expected %v and got %v \n", value, r.Value)
		}
	}
	fmt.Println("")
}

func TestInsert(t *testing.T) {
	tree := NewTree()

	//key := 1
	//value := []byte("test")
	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	key := datas.NewDataR(account)
	va1 := datas.NewDataAll(key)
	value := va1.Tobyte()
	var va datas.Data_All
	str := string(value)
	fmt.Println(str)
	err := json.Unmarshal(value, va)
	err = tree.Insert(*key, *va1)
	if err != nil {
		t.Errorf("%s", err)
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(va1) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}
}

func TestInsertSameKeyTwice(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)

	err := tree.Insert(*key, *value)
	if err != nil {
		t.Errorf("%s", err)
	}

	err = tree.Insert(*key, *datas.NewnullDataAll())
	if err == nil {
		t.Errorf("expected error but got nil")
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}

	if tree.Root.NumKeys > 1 {
		t.Errorf("expected 1 key and got %d", tree.Root.NumKeys)
	}
}

func TestInsertSameValueTwice(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)
	account1 := block.NewAcconunt(*block.NewAccountKey(2), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key1 := datas.NewDataR(account1)
	err := tree.Insert(*key, *value)
	if err != nil {
		t.Errorf("%s", err)
	}
	err = tree.Insert(*key1, *value)
	if err != nil {
		t.Errorf("%s", err)
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}

	if tree.Root.NumKeys <= 1 {
		t.Errorf("expected more than 1 key and got %d", tree.Root.NumKeys)
	}
}

func TestFindNilRoot(t *testing.T) {
	tree := NewTree()
	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	//value := datas.NewDataAll(key)
	r, err, _ := tree.Find(*key, false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	if r != nil {
		t.Errorf("expected nil got %s \n", r.Value.ToStr())
	}
}

func TestFind(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)

	err := tree.Insert(*key, *value)
	if err != nil {
		t.Errorf("%s", err)
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}
}

func TestDeleteNilTree(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	//value := datas.NewDataAll(key)

	err := tree.Delete(*key)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	r, err, _ := tree.Find(*key, false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	if r != nil {
		t.Errorf("returned struct after delete \n")
	}
}

func TestDelete(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)

	err := tree.Insert(*key, *value)
	if err != nil {
		t.Errorf("%s", err)
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}

	err = tree.Delete(*key)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	r, err, _ = tree.Find(*key, false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	if r != nil {
		t.Errorf("returned struct after delete \n")
	}
}

func TestDeleteNotFound(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)

	err := tree.Insert(*key, *value)
	if err != nil {
		t.Errorf("%s", err)
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}
	account1 := block.NewAcconunt(*block.NewAccountKey(2), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key1 := datas.NewDataR(account1)
	err = tree.Delete(*key1)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	r, err, _ = tree.Find(*key1, false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}
}

func TestMultiInsertSingleDelete(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)
	for i := 1; i < 6; i++ {
		//key := i
		//value := []byte("test")
		account = block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})

		key = datas.NewDataR(account)
		value = datas.NewDataAll(key)

		//fmt.Println(i)
		err := tree.Insert(*key, *value)
		if err != nil {
			t.Errorf("%s", err)
		}
	}

	r, err, _ := tree.Find(*key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}

	err = tree.Delete(*key)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	r, err, _ = tree.Find(*key, false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	if r != nil {
		t.Errorf("returned struct after delete - %v \n", r)
	}
}

func TestMultiInsertMultiDelete(t *testing.T) {
	tree := NewTree()

	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)
	var keys []datas.Data_R
	var values []*datas.Data_All
	for i := 1; i < 6; i++ {
		//key := i
		//value := []byte("test")
		account = block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})

		key = datas.NewDataR(account)
		value = datas.NewDataAll(key)

		//fmt.Println(i)
		keys = append(keys, *key)
		values = append(values, value)
		err := tree.Insert(*key, *value)
		if err != nil {
			t.Errorf("%s", err)
		}
	}

	r, err, _ := tree.Find(keys[0], false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, value) {
	if !r.Value.Equals(values[0]) {
		t.Errorf("expected %v and got %v \n", values[0], r.Value)
	}

	err = tree.Delete(keys[0])
	if err != nil {
		t.Errorf("%s\n", err)
	}

	r, err, _ = tree.Find(keys[0], false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	if r != nil {
		t.Errorf("returned struct after delete - %v \n", r)
	}

	r, err, _ = tree.Find(keys[3], false)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if r == nil {
		t.Errorf("returned nil \n")
	}

	//if !reflect.DeepEqual(r.Value, append(value, []byte("world3")...)) {
	if !r.Value.Equals(values[3]) {
		t.Errorf("expected %v and got %v \n", values[3], r.Value)
	}

	err = tree.Delete(keys[3])
	if err != nil {
		t.Errorf("%s\n", err)
	}

	r, err, _ = tree.Find(keys[3], false)
	if err == nil {
		t.Errorf("expected error and got nil")
	}

	if r != nil {
		t.Errorf("returned struct after delete - %v \n", r)
	}
}

func TestSign(t *testing.T) {
	tree := NewTree()
	for i := 100; i > 0; i-- {
		//key := i
		//value := []byte("test")
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)
		//	fmt.Println(i)
		err := tree.Insert(*key, *value)
		if err != nil {
			t.Errorf("%s", err)
		}

		r, err, _ := tree.Find(*key, false)
		if err != nil {
			t.Errorf("%s\n", err)
		}

		if r == nil {
			t.Errorf("returned nil \n")
		}

		//if !reflect.DeepEqual(r.Value, value) {
		//if !reflect.DeepEqual(r.Value, value) {
		if !r.Value.Equals(value) {
			t.Errorf("expected %v and got %v \n", value, r.Value)
		}
	}
	tree.SignAll()
	tree1 := NewTree()
	for i := 100; i > 0; i-- {
		//key := i
		//value := []byte("test")
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)
		//	fmt.Println(i)
		err := tree1.Insert(*key, *value)
		if err != nil {
			t.Errorf("%s", err)
		}

		r, err, _ := tree1.Find(*key, false)
		if err != nil {
			t.Errorf("%s\n", err)
		}

		if r == nil {
			t.Errorf("returned nil \n")
		}

		//if !reflect.DeepEqual(r.Value, value) {
		//if !reflect.DeepEqual(r.Value, value) {
		if !r.Value.Equals(value) {
			t.Errorf("expected %v and got %v \n", value, r.Value)
		}
	}
	tree1.SignAll()
	a := tree.Root.ShowSign()
	b := tree1.Root.ShowSign()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("expected %v and got %v \n", a, b)
	}

}

func TestVo(t *testing.T) {
	tree := NewTree()
	for i := 100; i > 0; i-- {
		//key := i
		//value := []byte("test")
		account := block.NewAcconunt(*block.NewAccountKey(i), 0, block.Pointers{})
		key := datas.NewDataR(account)
		value := datas.NewDataAll(key)
		//fmt.Println(i)
		err := tree.Insert(*key, *value)
		if err != nil {
			t.Errorf("%s", err)
		}

		r, err, _ := tree.Find(*key, false)
		if err != nil {
			t.Errorf("%s\n", err)
		}

		if r == nil {
			t.Errorf("returned nil \n")
		}

		//if !reflect.DeepEqual(r.Value, value) {
		//if !reflect.DeepEqual(r.Value, value) {
		if !r.Value.Equals(value) {
			t.Errorf("expected %v and got %v \n", value, r.Value)
		}
	}
	tree.SignAll()
	account := block.NewAcconunt(*block.NewAccountKey(9), 0, block.Pointers{})
	key := datas.NewDataR(account)
	r, error, node := tree.Find(*key, false)
	if error != nil {
		t.Errorf("%s\n", err)
	}
	vo := tree.MakeVo(r, node)
	ver := vo.Verifiable()
	if !ver {
		t.Errorf("vo is not true")
	}

}

func TestAll(t *testing.T) {
	account := block.NewAcconunt(*block.NewAccountKey(1), 0, block.Pointers{})
	//acc := account.(datas.Data_R)
	key := datas.NewDataR(account)
	value := datas.NewDataAll(key)

	tree := NewTree()

	err := tree.Insert(*key, *value)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	r, err, _ := tree.Find(*key, true)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	//fmt.Printf("%s\n\n", r.Value)
	fmt.Println(r.Value)

	tree.FindAndPrint(*key, true)
}
