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

type Node struct {
	Key    int
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
}

type rbTree struct {
	Root *Node
	size int
}

func (tree *rbTree) Insert(key int) {
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

func (tree *rbTree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *rbTree) insertCase2(node *Node) {

}
