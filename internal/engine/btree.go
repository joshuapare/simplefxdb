package engine

type BTreeNode struct {
	left  *BTreeNode
	right *BTreeNode
	data  int64
}

type BTree struct {
	root *BTreeNode
}

func (t *BTree) insert(data int64) *BTree {
	if t.root == nil {
		t.root = &BTreeNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BTreeNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BTreeNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BTreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}
