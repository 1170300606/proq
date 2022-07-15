package useless

import (
	"ProQueries/structs/merkletree/datas"
	"crypto/sha1"
	"fmt"
	"io"
)

const M = 4
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
const LIMIT_M_2 = (M + 1) / 2

type Position *BPlusFullNode

type BPlusTree struct {
	keyMax   int
	root     *BPlusFullNode
	ptr      *BPlusFullNode
	most     *datas.Data_All
	roothash []byte
	keynum   int
}

func MallocNewNode(isLeaf bool) *BPlusFullNode {
	var NewNode *BPlusFullNode
	if isLeaf == true {
		NewLeaf := MallocNewLeaf()
		NewNode = &BPlusFullNode{
			KeyNum:   0,
			Key:      make([]datas.Data_R, M+1), //申请M + 1是因为插入时可能暂时出现节点key大于M 的情况,待后期再分裂处理
			hash:     make([]datas.Data_S, M+1),
			isLeaf:   isLeaf,
			Children: nil,
			leafNode: NewLeaf,
		}
	} else {
		NewNode = &BPlusFullNode{
			KeyNum:   0,
			Key:      make([]datas.Data_R, M+1),
			hash:     make([]datas.Data_S, M+1),
			isLeaf:   isLeaf,
			Children: make([]*BPlusFullNode, M+1),
			leafNode: nil,
		}
	}
	for i, _ := range NewNode.Key {
		NewNode.Key[i] = *datas.NewnullDataR()
	}

	return NewNode
}

func MallocNewLeaf() *BPlusLeafNode {

	NewLeaf := BPlusLeafNode{
		Before: nil,
		Next:   nil,
		datas:  make([]*datas.Data_All, M+1),
	}
	for i, _ := range NewLeaf.datas {
		//NewLeaf.datas[i] = i
		NewLeaf.datas[i] = datas.NewnullDataAll()
	}
	return &NewLeaf
}

func (tree *BPlusTree) Initialize() {

	/* 根结点 */
	T := MallocNewNode(true)
	tree.ptr = T
	tree.root = T
	//ac := account.NewNullAcconunt()
	r := datas.NewnullDataR()
	tree.most = datas.NewDataAll(r)
	tree.keynum = 0
}

func (tree *BPlusTree) GetRootHash() []byte {
	return tree.roothash
}

//get before
func (tree *BPlusTree) getbf(i int, currentNode *BPlusFullNode) (int, *BPlusFullNode) {
	if i == 0 {
		if currentNode.leafNode.Before != nil {
			currentNode = currentNode.leafNode.Before
			i = currentNode.KeyNum - 1
		}
	} else {
		i--
	}
	return i, currentNode
}

//get after
func (tree *BPlusTree) getaft(i int, currentNode *BPlusFullNode) (int, *BPlusFullNode) {
	if i == currentNode.KeyNum {
		if currentNode.leafNode.Next != nil {
			currentNode = currentNode.leafNode.Next
			i = 0
		}
	} else {
		i++
	}

	return i, currentNode
}

func (tree *BPlusTree) rangeQUery(key datas.Data_R) (int, *BPlusFullNode) {
	var currentNode *BPlusFullNode
	var index int
	currentNode = tree.root

	for true {
		index = 0
		//for key >= currentNode.Key[index] && index < currentNode.KeyNum{
		for key.Ishigher(currentNode.Key[index]) && index < currentNode.KeyNum {
			index++
		}
		if index == 0 {
			if currentNode.isLeaf == false {
				currentNode = currentNode.Children[index]
			} else {
				return index, currentNode
			}

		} else {
			index--
			if currentNode.isLeaf == false {
				currentNode = currentNode.Children[index]
			} else {
				//if key == currentNode.Key[index] {
				if key.Equals(currentNode.Key[index]) {
					return index, currentNode
				} else {
					if key.Ishigher(currentNode.Key[index]) {
						return index, currentNode
					} else {
						return tree.getaft(index, currentNode)
					}
				}
			}
		}
	}
	return -1, currentNode
}

func (tree *BPlusTree) getvo(left datas.Data_R, right datas.Data_R) datas.Vo {

	vo1 := datas.MallocVo(tree.roothash)
	vo1.Root = tree.regetvo(left, right, vo1.Root, tree.root)
	vo1.Set(tree.roothash)
	return *vo1
}

func (tree *BPlusTree) regetvo(left datas.Data_R, right datas.Data_R, vonode *datas.Vonode, node *BPlusFullNode) *datas.Vonode {
	vonode = datas.MallocNewLeaf(node.KeyNum, node.isLeaf)
	i := 0
	if !node.isLeaf {

		for i = 0; i < (node.KeyNum - 1); i++ {
			if !(!node.Key[i].Islower(right) || !node.Key[i+1].Ishigher(left)) {
				vonode.Children[i] = tree.regetvo(left, right, vonode.Children[i], node.Children[i])
			} else {
				vonode.Children[i] = datas.MallocNewLeaf(0, true)
				vonode.Children[i].Sethash(node.hash[i].Showsign())
			}
		}
		if node.Key[i].Islower(right) {
			vonode.Children[i] = tree.regetvo(left, right, vonode.Children[i], node.Children[i])
		} else {
			vonode.Children[i] = datas.MallocNewLeaf(0, true)
			vonode.Children[i].Sethash(node.hash[i].Showsign())
		}
	} else {
		if !node.Key[0].Islower(right) || !node.Key[node.KeyNum-1].Ishigher(left) {
			vonode.Sethash(node.Getwholehash())
		}
		vonode.Setleaf(node.leafNode.Getdata())
	}
	return vonode
}

func (tree *BPlusTree) RangeQUery(left datas.Data_R, right datas.Data_R) datas.Vo {
	i, node_1 := tree.rangeQUery(left)
	i, node_1 = tree.getbf(i, node_1)
	j, node_2 := tree.rangeQUery(right)
	//left = *datas.NewDataR(*node_1.leafNode.datas[i].Showdata())
	//right = *datas.NewDataR(*node_2.leafNode.datas[j].Showdata())
	left = node_1.leafNode.datas[i].Showdata().Copy()
	right = node_2.leafNode.datas[j].Showdata().Copy()
	return tree.getvo(left, right)
}

func (tree *BPlusTree) Signall() {
	if tree.ptr == nil {
		return
	}
	tree.root.Hash1()
	h := sha1.New()
	for i := 0; i < tree.root.KeyNum; i++ {
		io.WriteString(h, string(tree.root.hash[i].Showsign()))
	}
	tree.roothash = h.Sum(nil)
}

func FindMostLeft(P Position) Position {
	var Tmp Position
	Tmp = P
	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.Children[0].isLeaf == true {
		return Tmp.Children[0]
	} else {
		for Tmp != nil && Tmp.Children[0].isLeaf != true {
			Tmp = Tmp.Children[0]
		}
	}
	return Tmp.Children[0]
}

func FindMostRight(P Position) Position {
	var Tmp Position
	Tmp = P

	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.Children[Tmp.KeyNum-1].isLeaf == true {
		return Tmp.Children[Tmp.KeyNum-1]
	} else {
		for Tmp != nil && Tmp.Children[Tmp.KeyNum-1].isLeaf != true {
			Tmp = Tmp.Children[Tmp.KeyNum-1]
		}
	}

	return Tmp.Children[Tmp.KeyNum-1]
}

/* 寻找一个兄弟节点，其存储的关键字未满，若左右都满返回nil */
func FindSibling(Parent Position, i int) Position {
	var Sibling Position
	var upperLimit int
	upperLimit = M
	Sibling = nil
	if i == 0 {
		if Parent.Children[1].KeyNum < upperLimit {

			Sibling = Parent.Children[1]
		}
	} else if Parent.Children[i-1].KeyNum < upperLimit {
		Sibling = Parent.Children[i-1]
	} else if i+1 < Parent.KeyNum && Parent.Children[i+1].KeyNum < upperLimit {
		Sibling = Parent.Children[i+1]
	}
	return Sibling
}

/* 查找兄弟节点，其关键字数大于M/2 ;没有返回nil j用来标识是左兄还是右兄*/

func FindSiblingKeyNum_M_2(Parent Position, i int, j *int) Position {
	var lowerLimit int
	var Sibling Position
	Sibling = nil

	lowerLimit = LIMIT_M_2

	if i == 0 {
		if Parent.Children[1].KeyNum > lowerLimit {
			Sibling = Parent.Children[1]
			*j = 1
		}
	} else {
		if Parent.Children[i-1].KeyNum > lowerLimit {
			Sibling = Parent.Children[i-1]
			*j = i - 1
		} else if i+1 < Parent.KeyNum && Parent.Children[i+1].KeyNum > lowerLimit {
			Sibling = Parent.Children[i+1]
			*j = i + 1
		}

	}
	return Sibling
}

/* 当要对X插入data的时候，i是X在Parent的位置，insertIndex是data要插入的位置，j可由查找得到
   当要对Parent插入X节点的时候，posAtParent是要插入的位置，Key和j的值没有用
*/
func (tree *BPlusTree) InsertElement(isData bool, Parent Position, X Position, Key datas.Data_R, posAtParent int, insertIndex int, data *datas.Data_All) Position {

	var k int
	if isData {
		/* 插入data*/
		k = X.KeyNum - 1
		for k >= insertIndex {
			X.Key[k+1] = X.Key[k]
			X.leafNode.datas[k+1] = X.leafNode.datas[k]
			k--
		}

		X.Key[insertIndex] = Key
		X.leafNode.datas[insertIndex] = data
		if Parent != nil {
			Parent.Key[posAtParent] = X.Key[0] //可能min_key 已发生改变
		}

		X.KeyNum++

	} else {
		/* 插入节点 */
		/* 对树叶节点进行连接 */
		if X.isLeaf == true {
			if posAtParent > 0 {
				Parent.Children[posAtParent-1].leafNode.Next = X //TODO do Before
				X.leafNode.Before = Parent.Children[posAtParent-1]
			}
			X.leafNode.Next = Parent.Children[posAtParent] //TODO do Before
			if Parent.Children[posAtParent] != nil {
				Parent.Children[posAtParent].leafNode.Before = X
			}
			//更新叶子指针
			//if X.Key[0] <= tree.ptr.Key[0] {
			if X.Key[0].Islower(tree.ptr.Key[0]) {
				tree.ptr = X
			}
		}

		k = Parent.KeyNum - 1
		for k >= posAtParent { //插入节点时key也要对应的插入
			Parent.Children[k+1] = Parent.Children[k]
			Parent.Key[k+1] = Parent.Key[k]
			k--
		}
		Parent.Key[posAtParent] = X.Key[0]
		Parent.Children[posAtParent] = X
		Parent.KeyNum++
	}

	return X
}

/*
	两个参数X posAtParent 有些重复 posAtParent可以通过X的最小关键字查找得到
*/
func (tree *BPlusTree) RemoveElement(isData bool, Parent Position, X Position, posAtParent int, deleteIndex int) Position {

	var k, keyNum int

	if isData {
		keyNum = X.KeyNum
		/* 删除key */
		k = deleteIndex + 1
		for k < keyNum {
			X.Key[k-1] = X.Key[k]
			X.leafNode.datas[k-1] = X.leafNode.datas[k]
			k++
		}

		X.Key[keyNum-1] = *datas.NewnullDataR()
		X.leafNode.datas[keyNum-1] = datas.NewnullDataAll()
		Parent.Key[posAtParent] = X.Key[0]
		X.KeyNum--
	} else {
		/* 删除节点 */
		/* 修改树叶节点的链接 */
		if X.isLeaf == true && posAtParent > 0 {
			Parent.Children[posAtParent-1].leafNode.Next = Parent.Children[posAtParent+1] //TODO Before
			Parent.Children[posAtParent+1].leafNode.Before = Parent.Children[posAtParent-1]
		}

		keyNum = Parent.KeyNum
		k = posAtParent + 1
		for k < keyNum {
			Parent.Children[k-1] = Parent.Children[k]
			Parent.Key[k-1] = Parent.Key[k]
			k++
		}

		//if X.Key[0] == tree.ptr.Key[0] { // refresh ptr
		if X.Key[0].Equals(tree.ptr.Key[0]) {
			tree.ptr = Parent.Children[0]
		}
		Parent.Children[Parent.KeyNum-1] = nil
		Parent.Key[Parent.KeyNum-1] = *datas.NewnullDataR()

		Parent.KeyNum--

	}
	return X
}

/* Src和Dst是两个相邻的节点，posAtParent是Src在Parent中的位置；
将Src的元素移动到Dst中 ,eNum是移动元素的个数*/
func (tree *BPlusTree) MoveElement(src Position, dst Position, parent Position, posAtParent int, eNum int) Position {
	//var TmpKey, data int
	var TmpKey datas.Data_R
	var data *datas.Data_All
	var Child Position
	var j int
	var srcInFront bool

	srcInFront = false

	//if src.Key[0] < dst.Key[0] {
	if !src.Key[0].Ishigher(dst.Key[0]) {
		srcInFront = true
	}
	j = 0
	/* 节点Src在Dst前面 */
	if srcInFront {
		if src.isLeaf == false {
			for j < eNum {
				Child = src.Children[src.KeyNum-1]
				tree.RemoveElement(false, src, Child, src.KeyNum-1, INT_MIN) //每删除一个节点keyNum也自动减少1 队尾删
				//tree.InsertElement(false, dst, Child, INT_MIN, 0, INT_MIN, INT_MIN) //队头加
				tree.InsertElement(false, dst, Child, *datas.NewnullDataR(), 0, INT_MIN, datas.NewnullDataAll()) //队头加
				j++
			}
		} else {
			for j < eNum {
				TmpKey = src.Key[src.KeyNum-1]
				data = src.leafNode.datas[src.KeyNum-1]
				tree.RemoveElement(true, parent, src, posAtParent, src.KeyNum-1)
				tree.InsertElement(true, parent, dst, TmpKey, posAtParent+1, 0, data)
				j++
			}

		}

		parent.Key[posAtParent+1] = dst.Key[0]
		/* 将树叶节点重新连接 */
		if src.KeyNum > 0 {
			FindMostRight(src).leafNode.Next = FindMostLeft(dst) //TODO Before 似乎不需要重连，src的最右本身就是dst最左的上一元素
			FindMostLeft(dst).leafNode.Before = FindMostRight(src)
		} else {
			if src.isLeaf == true {
				parent.Children[posAtParent-1].leafNode.Next = dst //TODO Before
				dst.leafNode.Before = parent.Children[posAtParent-1]
			}
			//  此种情况肯定是merge merge中有实现先移动再删除操作
			//tree.RemoveElement(false ,parent.parent，parent ,parentIndex,INT_MIN )
		}
	} else {
		if src.isLeaf == false {
			for j < eNum {
				Child = src.Children[0]
				tree.RemoveElement(false, src, Child, 0, INT_MIN) //从src的队头删
				tree.InsertElement(false, dst, Child, *datas.NewnullDataR(), dst.KeyNum, INT_MIN, datas.NewnullDataAll())
				j++
			}

		} else {
			for j < eNum {
				TmpKey = src.Key[0]
				data = src.leafNode.datas[0]
				tree.RemoveElement(true, parent, src, posAtParent, 0)
				tree.InsertElement(true, parent, dst, TmpKey, posAtParent-1, dst.KeyNum, data)
				j++
			}

		}

		parent.Key[posAtParent] = src.Key[0]
		if src.KeyNum > 0 {
			FindMostRight(dst).leafNode.Next = FindMostLeft(src) //TODO before
			FindMostLeft(src).leafNode.Before = FindMostRight(dst)
		} else {
			if src.isLeaf == true {
				dst.leafNode.Next = src.leafNode.Next //TODO before
				src.leafNode.Next.leafNode.Before = dst
			}
			//tree.RemoveElement(false ,parent.parent，parent ,parentIndex,INT_MIN )
		}
	}

	return parent
}

//i为节点X的位置
func (tree *BPlusTree) SplitNode(Parent Position, beSplitedNode Position, i int) Position {
	var j, k, keyNum int
	var NewNode Position

	if beSplitedNode.isLeaf == true {
		NewNode = MallocNewNode(true)
	} else {
		NewNode = MallocNewNode(false)
	}

	k = 0
	j = beSplitedNode.KeyNum / 2
	keyNum = beSplitedNode.KeyNum
	for j < keyNum {
		if beSplitedNode.isLeaf == false { //Internal node
			NewNode.Children[k] = beSplitedNode.Children[j]
			beSplitedNode.Children[j] = nil
		} else {
			NewNode.leafNode.datas[k] = beSplitedNode.leafNode.datas[j]
			beSplitedNode.leafNode.datas[j] = datas.NewnullDataAll()
		}
		NewNode.Key[k] = beSplitedNode.Key[j]
		beSplitedNode.Key[j] = *datas.NewnullDataR()
		NewNode.KeyNum++
		beSplitedNode.KeyNum--
		j++
		k++
	}

	if Parent != nil {
		tree.InsertElement(false, Parent, NewNode, *datas.NewnullDataR(), i+1, INT_MIN, datas.NewnullDataAll())
		// parent > limit 时的递归split recurvie中实现
	} else {
		/* 如果是X是根，那么创建新的根并返回 */
		Parent = MallocNewNode(false)
		tree.InsertElement(false, Parent, beSplitedNode, *datas.NewnullDataR(), 0, INT_MIN, datas.NewnullDataAll())
		tree.InsertElement(false, Parent, NewNode, *datas.NewnullDataR(), 1, INT_MIN, datas.NewnullDataAll())
		tree.root = Parent
		return Parent
	}

	return beSplitedNode
	// 为什么返回一个X一个Parent?
}

/* 合并节点,X少于M/2关键字，S有大于或等于M/2个关键字*/
func (tree *BPlusTree) MergeNode(Parent Position, X Position, S Position, i int) Position {
	var Limit int

	/* S的关键字数目大于M/2 */
	if S.KeyNum > LIMIT_M_2 {
		/* 从S中移动一个元素到X中 */
		tree.MoveElement(S, X, Parent, i, 1)
	} else {
		/* 将X全部元素移动到S中，并把X删除 */
		Limit = X.KeyNum
		tree.MoveElement(X, S, Parent, i, Limit) //最多时S恰好MAX MoveElement已考虑了parent.key的索引更新
		tree.RemoveElement(false, Parent, X, i, INT_MIN)
	}
	return Parent
}

func (tree *BPlusTree) RecursiveInsert(beInsertedElement Position, Key datas.Data_R, posAtParent int, Parent Position, data *datas.Data_All) (Position, bool) {
	var InsertIndex, upperLimit int
	var Sibling Position
	var result bool
	result = true
	/* 查找分支 */
	InsertIndex = 0
	//for InsertIndex < beInsertedElement.KeyNum && Key >= beInsertedElement.Key[InsertIndex] {
	for InsertIndex < beInsertedElement.KeyNum && Key.Ishigher(beInsertedElement.Key[InsertIndex]) {
		/* 重复值不插入 */
		//if Key == beInsertedElement.Key[InsertIndex] {
		if Key.Equals(beInsertedElement.Key[InsertIndex]) {
			return beInsertedElement, false
		}
		InsertIndex++
	}
	//key必须大于被插入节点的最小元素，才能插入到此节点，故需回退一步
	if InsertIndex != 0 && beInsertedElement.isLeaf == false {
		InsertIndex--
	}

	/* 树叶 */
	if beInsertedElement.isLeaf == true {
		beInsertedElement = tree.InsertElement(true, Parent, beInsertedElement, Key, posAtParent, InsertIndex, data) //返回叶子节点
		/* 内部节点 */
	} else {
		beInsertedElement.Children[InsertIndex], result = tree.RecursiveInsert(beInsertedElement.Children[InsertIndex], Key, InsertIndex, beInsertedElement, data)
		//更新parent发生在split时
	}
	/* 调整节点 */

	upperLimit = M
	if beInsertedElement.KeyNum > upperLimit {
		/* 根 */
		if Parent == nil {
			/* 分裂节点 */
			beInsertedElement = tree.SplitNode(Parent, beInsertedElement, posAtParent)
		} else {
			Sibling = FindSibling(Parent, posAtParent)
			if Sibling != nil {
				/* 将T的一个元素（Key或者Child）移动的Sibing中 */
				tree.MoveElement(beInsertedElement, Sibling, Parent, posAtParent, 1)
			} else {
				/* 分裂节点 */
				beInsertedElement = tree.SplitNode(Parent, beInsertedElement, posAtParent)
			}
		}

	}
	if Parent != nil {
		Parent.Key[posAtParent] = beInsertedElement.Key[0]
	}

	return beInsertedElement, result
}

/* 插入 */
func (tree *BPlusTree) Insert(Key datas.Data_R, data *datas.Data_All) (Position, bool) {

	po, bo := tree.RecursiveInsert(tree.root, Key, 0, nil, data) //从根节点开始插入
	if bo {
		tree.keynum++
	}
	return po, bo
}

func (tree *BPlusTree) RecursiveRemove(beRemovedElement Position, Key datas.Data_R, posAtParent int, Parent Position) (Position, bool) {

	var deleteIndex int
	var Sibling Position
	var NeedAdjust bool
	var result bool
	Sibling = nil

	/* 查找分支   TODO查找函数可以在参考这里的代码 或者实现一个递归遍历*/
	deleteIndex = 0
	//for deleteIndex < beRemovedElement.KeyNum && Key >= beRemovedElement.Key[deleteIndex] {
	for deleteIndex < beRemovedElement.KeyNum && Key.Ishigher(beRemovedElement.Key[deleteIndex]) {
		//if Key == beRemovedElement.Key[deleteIndex] {
		if Key.Equals(beRemovedElement.Key[deleteIndex]) {
			break
		}
		deleteIndex++
	}

	if beRemovedElement.isLeaf == true {
		/* 没找到 */
		//if Key != beRemovedElement.Key[deleteIndex] || deleteIndex == beRemovedElement.KeyNum {
		if !Key.Equals(beRemovedElement.Key[deleteIndex]) || deleteIndex == beRemovedElement.KeyNum {
			return beRemovedElement, false
		}
	} else {
		//if deleteIndex == beRemovedElement.KeyNum || Key < beRemovedElement.Key[deleteIndex] {
		if deleteIndex == beRemovedElement.KeyNum || !Key.Ishigher(beRemovedElement.Key[deleteIndex]) {
			deleteIndex-- //准备到下层节点查找
		}
	}

	/* 树叶 */
	if beRemovedElement.isLeaf == true {
		beRemovedElement = tree.RemoveElement(true, Parent, beRemovedElement, posAtParent, deleteIndex)
	} else {
		beRemovedElement.Children[deleteIndex], result = tree.RecursiveRemove(beRemovedElement.Children[deleteIndex], Key, deleteIndex, beRemovedElement)
	}

	NeedAdjust = false
	//有子节点的root节点，当keyNum小于2时
	if Parent == nil && beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < 2 {
		NeedAdjust = true
	} else if Parent != nil && beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < LIMIT_M_2 {
		/* 除根外，所有中间节点的儿子数不在[M/2]到M之间时。(符号[]表示向上取整) */
		NeedAdjust = true
	} else if Parent != nil && beRemovedElement.isLeaf == true && beRemovedElement.KeyNum < LIMIT_M_2 {
		/* （非根）树叶中关键字的个数不在[M/2]到M之间时 */
		NeedAdjust = true
	}

	/* 调整节点 */
	if NeedAdjust {
		/* 根 */
		if Parent == nil {
			if beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < 2 {
				//树根的更新操作 树高度减一
				beRemovedElement = beRemovedElement.Children[0]
				tree.root = beRemovedElement.Children[0]
				return beRemovedElement, true
			}

		} else {
			/* 查找兄弟节点，其关键字数目大于M/2 */
			Sibling = FindSiblingKeyNum_M_2(Parent, posAtParent, &deleteIndex)
			if Sibling != nil {
				tree.MoveElement(Sibling, beRemovedElement, Parent, deleteIndex, 1)
			} else {
				if posAtParent == 0 {
					Sibling = Parent.Children[1]
				} else {
					Sibling = Parent.Children[posAtParent-1]
				}

				Parent = tree.MergeNode(Parent, beRemovedElement, Sibling, posAtParent)
				//Merge中已考虑空节点的删除
				beRemovedElement = Parent.Children[posAtParent]
			}
		}

	}

	return beRemovedElement, result
}

/* 删除 */
func (tree *BPlusTree) Remove(Key datas.Data_R) (Position, bool) {
	return tree.RecursiveRemove(tree.root, Key, 0, nil)
}

func (tree *BPlusTree) FindData(key datas.Data_R) (*datas.Data_All, bool) {
	var currentNode *BPlusFullNode
	var index int
	currentNode = tree.root
	for index < currentNode.KeyNum {
		index = 0
		//for key >= currentNode.Key[index] && index < currentNode.KeyNum{
		//for key.Ishigher(currentNode.Key[index]) && index < currentNode.KeyNum {
		for index < currentNode.KeyNum && key.Ishigher(currentNode.Key[index]) {
			index++
		}
		if index == 0 {
			return datas.NewnullDataAll(), false
		} else {
			index--
			if currentNode.isLeaf == false {
				currentNode = currentNode.Children[index]
				index = 0
			} else {
				//if key == currentNode.Key[index] {
				if key.Equals(currentNode.Key[index]) {
					return currentNode.leafNode.datas[index], true
				} else {
					return datas.NewnullDataAll(), false
				}
			}
		}

	}
	return datas.NewnullDataAll(), false
}

func (tree *BPlusTree) ShowAll() {
	i := 0
	var tmp *BPlusFullNode
	tmp = tree.ptr
	for tmp != nil {
		i = 0
		for i < tmp.KeyNum {
			tmp.leafNode.datas[i].Showdata()
			fmt.Println(tmp.leafNode.datas[i].Showdata())
			i++
		}
		tmp = tmp.leafNode.Next
	}
}
