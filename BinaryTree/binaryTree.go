package BinaryTree

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinTree struct {
	Root *Node
	size int
}

func NewBinTree() *BinTree {
	return &BinTree{Root: nil, size: 0}
}

// Insert functions
