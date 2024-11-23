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

// PreOrderTravers function performs a pre-order traversal of the binary tree.
func (tree *Tree) PreOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += strconv.Itoa(node.Key) + " "
		str += tree.PreOrderTravers(node.Left) + " "
		str += tree.PreOrderTravers(node.Right) + " "
	}
	return str
}
