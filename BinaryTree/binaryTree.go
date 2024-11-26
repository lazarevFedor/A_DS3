package BinaryTree

import (
	"fmt"
	"strconv"
)

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

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func Output(node *Node, prefix string, isTail bool, str *string) {
	if str == nil || node == nil {
		return
	}
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "    "
		}
		Output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "    "
		}
		Output(node.Left, newPrefix, true, str)
	}
}
