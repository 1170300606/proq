package useless

import (
	"ProQueries/structs/merkletree/datas"
	"crypto/sha1"
	"io"
)

type BPlusFullNode struct {
	KeyNum   int
	Key      []datas.Data_R
	hash     []datas.Data_S
	isLeaf   bool
	Children []*BPlusFullNode
	leafNode *BPlusLeafNode
}

func (node *BPlusFullNode) Hash1() []datas.Data_S {
	if node.hash[0].Showsign() != nil {
		return node.hash
	}
	h := sha1.New()
	if !node.isLeaf {
		for i := 0; i < node.KeyNum; i++ {
			hashs := node.Children[i].Hash1()
			for j := 0; j < len(hashs); j++ {
				io.WriteString(h, string(hashs[j].Showsign()))
			}
			node.hash[i] = *datas.NewdataS(h.Sum(nil))
			h.Reset()
		}
	} else {
		for i := 0; i < node.KeyNum; i++ {
			//str := string(node.leafNode.datas[i].Showsign())
			//io.WriteString(h, str)
			node.hash[i] = *datas.NewdataS(node.leafNode.datas[i].Showsign())
			h.Reset()
		}
	}

	return node.hash
}

func (node *BPlusFullNode) Getwholehash() []byte {
	h := sha1.New()
	for i := 0; i < node.KeyNum; i++ {
		io.WriteString(h, string(node.hash[i].Showsign()))
		//a := node.hash[i].Showsign()
		//fmt.Println(a)
		//a = node.hash[i].Showsign()
		//fmt.Println(a)
	}
	a1 := h.Sum(nil)
	return a1
}

type BPlusLeafNode struct {
	Before *BPlusFullNode
	Next   *BPlusFullNode
	//hash   []datas.Data_S
	datas []*datas.Data_All
}

func (node *BPlusLeafNode) Getdata() []*datas.Data_All {
	return node.datas
}
