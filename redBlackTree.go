package main

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
	Color  string
}

func newNode(key int) *Node {
	return &Node{Key: key, Left: nullNode, Right: nullNode, Parent: nullNode, Color: "red"}
}

func nodeExists(n *Node) bool {
	if n != nullNode && n != nil {
		return true
	}
	return false
}

type Tree struct {
	Root *Node
}

var nullNode *Node = &Node{Key: 0, Left: nil, Right: nil, Parent: nil, Color: "black"}

func (t *Tree) rotateLeft(node *Node) *Node {
	var rightSubTree *Node = node.Right
	var rightLeftSubTree *Node = rightSubTree.Left
	rightSubTree.Left = node
	node.Right = rightLeftSubTree
	node = rightSubTree
	return node
}

func (t *Tree) rotateRight(node *Node) *Node {
	var leftSubTree *Node = node.Left
	var leftRightSubTree *Node = leftSubTree.Right
	leftSubTree.Right = node
	node.Left = leftRightSubTree
	node = leftSubTree
	return node
}

func (t *Tree) balanceTree(node *Node) *Node {
	var uncle *Node
	for node.Parent != nullNode && node.Parent.Color == "red" {
		if node.Parent == node.Parent.Parent.Left {
			uncle = node.Parent.Parent.Right
			if uncle.Color == "red" {
				node.Parent.Color = "black"
				uncle.Color = "black"
				node.Parent.Parent.Color = "red"
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					t.rotateLeft(node)
				}
				node.Parent.Color = "black"
				node.Parent.Parent.Color = "red"
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle = node.Parent.Parent.Left
			if uncle.Color == "red" {
				node.Parent.Color = "black"
				uncle.Color = "black"
				node.Parent.Parent.Color = "red"
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					t.rotateRight(node)
				}
				node.Parent.Color = "black"
				node.Parent.Parent.Color = "red"
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = "black"
	return node
}

func (t *Tree) Insert(key int) {
	var current *Node = t.Root
	var parent *Node = nullNode
	for nodeExists(current) {
		parent = current
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	var node *Node = newNode(key)
	node.Parent = parent
	if parent == nullNode {
		t.Root = node
	} else if key < parent.Key {
		parent.Left = node
	} else {
		parent.Right = node
	}
	t.balanceTree(node)
}
