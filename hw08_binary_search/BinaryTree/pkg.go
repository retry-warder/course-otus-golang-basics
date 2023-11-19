package binarytree

type Node struct {
	left  *Node
	right *Node
	value int
	data  string
}

type Bt struct {
	root *Node
}

func (bt *Bt) Root() *Node {
	return bt.root
}

func NewBst() *Bt {
	bst := new(Bt)
	return bst
}

func (root *Node) insert(newnode *Node) {
	if newnode.value == root.value {
		root.data = newnode.data
		return
	} else if newnode.value > root.value {
		if root.right == nil {
			root.right = newnode
		} else {
			root.right.insert(newnode)
		}
		return
	}
	if newnode.value < root.value {
		if root.left == nil {
			root.left = newnode
		} else {
			root.left.insert(newnode)
		}
		return
	}
}

func (bt *Bt) Insert(value int, data string) {
	if bt.root == nil {
		bt.root = &Node{nil, nil, value, data}
	}
	bt.root.insert(&Node{nil, nil, value, data})
}

func BSearch(root *Node, value int) (bool, string) {
	if root == nil {
		return false, ""
	}
	if value != root.value {
		if value < root.value {
			return BSearch(root.left, value)
		}
		return BSearch(root.right, value)
	}
	return true, root.data
}
