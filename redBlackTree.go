package main

type color bool

const (
	black, red color = true, false
)

func Comparator(value1, value2 int) int {
	if value1 < value2 {
		return -1
	}
	if value1 > value2 {
		return 1
	}
	return 0
}

// Node and Tree structs
type Node struct {
	Key    int
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
}

type RBTree struct {
	Root *Node
	size int
}

// Color of node
func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}

// Relatives of nodes
func (node *Node) grandParent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

// Insert and its cases
func (tree *RBTree) Insert(key int) {
	var insertedNode *Node
	if tree.Root == nil {
		tree.Root = &Node{Key: key, color: red}
		insertedNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := Comparator(key, node.Key)
			switch compare {
			case 0:
				node.Key = key
				return
			case -1:
				if node.Left == nil {
					node.Left = &Node{Key: key, color: red}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case 1:
				if node.Right == nil {
					node.Right = &Node{Key: key, color: red}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}

func (tree *RBTree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RBTree) insertCase2(node *Node) {
	if nodeColor(node.Parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *RBTree) insertCase3(node *Node) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandParent().color = red
		tree.insertCase1(node.grandParent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RBTree) insertCase4(node *Node) {
	grandparent := node.grandParent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *RBTree) insertCase5(node *Node) {
	node.Parent.color = black
	grandparent := node.grandParent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

// Rotates the tree
func (tree *RBTree) rotateLeft(node *Node) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *RBTree) rotateRight(node *Node) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *RBTree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}
