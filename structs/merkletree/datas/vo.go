package datas

import (
	"crypto/sha1"
	"crypto/sha256"
	"io"
)

type Vo struct {
	hash []byte
	Root *Vonode
}

func (vo *Vo) Set(hash []byte) {
	vo.hash = hash
}

type Vonode struct {
	hash     []byte
	KeyNum   int
	isLeaf   bool
	Children []*Vonode
	leafNode []*Data_All
}

func MallocVo(hash []byte) *Vo {
	vo := &Vo{
		hash: hash,
		Root: &Vonode{},
	}
	return vo
}

func MallocNewLeaf(keynum int, isleaf bool) *Vonode {
	NewNode := &Vonode{
		hash:     nil,
		KeyNum:   keynum,
		isLeaf:   isleaf,
		Children: make([]*Vonode, keynum),
		leafNode: nil,
	}
	return NewNode
}
func (vo *Vonode) Set(keynum int, isleaf bool) {
	vo.KeyNum = keynum
	vo.isLeaf = isleaf
	vo.Children = make([]*Vonode, keynum)
}

func (vo *Vonode) Sethash(hash []byte) {
	vo.hash = hash
}

func (vo *Vonode) Setleaf(data []*Data_All) {
	vo.leafNode = data
}

func (vo *Vonode) Gethash() []byte {
	h := sha256.New()
	if vo.hash == nil {

		if vo.leafNode == nil {
			for i := 0; i < vo.KeyNum; i++ {
				io.WriteString(h, string(vo.Children[i].Gethash()))
			}
			a2 := h.Sum(nil)
			//return h.Sum(nil)
			return a2
		} else {
			h = sha1.New()
			for i := 0; i < vo.KeyNum; i++ {
				//s1 := string(vo.leafNode[i].Showsign())
				io.WriteString(h, string(vo.leafNode[i].Showsign()))
				//a1 := h.Sum(nil)
				//fmt.Println(vo.leafNode[i].Showsign())

			}
			a1 := h.Sum(nil)
			return a1
		}

	}
	return vo.hash

}

func Newnullvo() *Vo {
	data_r := &Vo{}
	return data_r
}

func (vo *Vo) Verifiable() bool {

	hash := vo.Root.Gethash()
	s1 := string(vo.hash)
	s2 := string(hash)
	if s1 == s2 {
		return true
	}
	return false
}
