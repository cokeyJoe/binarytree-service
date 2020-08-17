// Package binary contains binary tree implementation
package binary

import "sync"

// Tree binary tree struct
type Tree struct {
	head *TreeNode

	mutex *sync.Mutex
}

type TreeNode struct {
	value int

	left  *TreeNode
	right *TreeNode
}

// New returns new empty binary tree
func New() *Tree {
	mutex := &sync.Mutex{}
	return &Tree{mutex: mutex}
}

// NewFromInts returns new binary tree with nodes with values
func NewFromInts(values ...int) *Tree {

	t := New()

	for _, v := range values {
		t.Insert(v)
	}

	return t
}

// Find searches for value in tree and returns value's Node if found
func (t *Tree) Find(value int) *TreeNode {
	if t.head == nil {
		return nil
	}

	return t.head.find(value)
}

func (tn *TreeNode) find(value int) *TreeNode {
	if tn.value == value {
		return tn
	}

	if value <= tn.value {

		if tn.left == nil {
			return nil
		}

		return tn.left.find(value)
	} else {
		if tn.right == nil {
			return nil
		}

		return tn.right.find(value)
	}
}

// Remove removes value if find and rebuilds tree
func (t *Tree) Remove(value int) {
	panic("Tree.Remove(value) not implemented")
}

// Insert insert value in tree
func (t *Tree) Insert(value int) {

	t.mutex.Lock()

	defer t.mutex.Unlock()

	if t.head == nil {

		t.head = &TreeNode{value: value, left: nil, right: nil}
	} else {

		if t.head.value == value {
			return
		}

		t.head.insert(value)

	}

}

func (tn *TreeNode) insert(value int) {

	if tn.value == value {
		return
	}

	if value <= tn.value {
		if tn.left == nil {
			tn.left = &TreeNode{value: value, left: nil, right: nil}
		} else {
			tn.left.insert(value)
		}
	} else {
		if tn.right == nil {
			tn.right = &TreeNode{value: value, left: nil, right: nil}
		} else {
			tn.right.insert(value)
		}
	}
}
