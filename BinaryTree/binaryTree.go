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

func (tree *Tree) PreOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += strconv.Itoa(node.Key) + " "
		str += tree.PreOrderTravers(node.Left) + " "
		str += tree.PreOrderTravers(node.Right) + " "
	}
	return str
}

func (tree *Tree) InOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += tree.InOrderTravers(node.Left) + " "
		str += strconv.Itoa(node.Key) + " "
		str += tree.InOrderTravers(node.Right) + " "
	}
	return str
}

func (tree *Tree) PostOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += tree.PostOrderTravers(node.Left) + " "
		str += tree.PostOrderTravers(node.Right) + " "
		str += strconv.Itoa(node.Key) + " "
	}
	return str
}

func (tree *Tree) LevelOrderTravers(root *Node) string {
	var str string
	if root == nil {
		return str
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		str += strconv.Itoa(node.Key) + " "
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return str
}
