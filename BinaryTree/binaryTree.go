package BinaryTree

import "strconv"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
	size int
}

func NewBinTree() *Tree {
	return &Tree{Root: nil, size: 0}
}

func (tree *Tree) preOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += strconv.Itoa(node.Key)
		str += tree.preOrderTravers(node.Left)
		str += tree.preOrderTravers(node.Right)
	}
	return str
}

func (tree *Tree) inOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += tree.inOrderTravers(node.Left)
		str += strconv.Itoa(node.Key)
		str += tree.inOrderTravers(node.Right)
	}
	return str
}

func (tree *Tree) postOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += tree.postOrderTravers(node.Left)
		str += tree.postOrderTravers(node.Right)
		str += strconv.Itoa(node.Key)
	}
	return str
}
