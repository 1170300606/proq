package merkletree

import (
	"crypto/sha256"
	"reflect"
)

type Vo2 struct {
	hash []byte
	Root *Vonode2
	Goal *Vonode2
	Po   int
}

type Vonode2 struct {
	hash   []byte
	KeyNum int
	//isLeaf   bool
	Children []*Vonode2
	leafNode [][]byte
}

func MallocVo(hash []byte, node1 *Vonode2, node2 *Vonode2, po int) *Vo2 {
	vo := &Vo2{
		hash: hash,
		//Root: &Vonode2{},
		//Goal: &Vonode2{},
		Root: node1,
		Goal: node2,
		Po:   po,
	}
	return vo
}

func MallocNewVonode2(keynum int, isleaf bool) *Vonode2 {
	NewNode := &Vonode2{
		hash:   nil,
		KeyNum: keynum,
		//isLeaf:   isleaf,
		Children: make([]*Vonode2, keynum),
		leafNode: nil,
	}
	return NewNode
}

func (node *Vonode2) Sethash(hash []byte) {
	node.hash = hash
}

func (node *Vonode2) Setleafs(hashs [][]byte) {
	node.leafNode = hashs
}

func (node *Vonode2) SetChildren(children []*Vonode2) {
	node.Children = children
}

func (node *Vonode2) Gethash() []byte {
	return node.hash
}

func (vo *Vo2) Verifiable() bool {
	var node *Vonode2
	//node_temp := &Vonode2{}
	node = vo.Root
	for node != vo.Goal {
		//node_temp = nil
		//h := sha256.New()
		if !node.Ver() { //判断中间节点的vonode
			return false
		}
		for i := 0; i < node.KeyNum; i++ {
			if node.Children[i].KeyNum != 0 {
				node = node.Children[i]
				break
			}
		}
	}
	if !node.Ver() { //判断最后包含叶子节点的vonode
		return false
	}
	return true

}

func (vo *Vonode2) Ver() bool {
	h := sha256.New()
	var a []byte
	if vo.leafNode != nil {
		for i := 0; i < vo.KeyNum; i++ {
			h.Write(vo.leafNode[i])
		}
		a = h.Sum(nil)
	} else {
		for i := 0; i < vo.KeyNum; i++ {
			h.Write(vo.Children[i].hash)
		}
		a = h.Sum(nil)
	}
	if reflect.DeepEqual(a, vo.hash) {
		return true
	}
	return false
}
