package engine

import (
	"fmt"
	"io"
	"os"
)

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

func print(w io.Writer, node *BTreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
}
