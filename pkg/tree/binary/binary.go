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

	mutex *sync.Mutex
}

// New returns new empty binary tree
func New() *Tree {
	mutex := &sync.Mutex{}
	return &Tree{mutex: mutex}
}

// NewFromInts returns new binary tree with nodes with values
func NewFromInts(values ...int) *Tree {
	panic("NewFromInts not implemented")
}

// Find searches for value in tree and returns value's Node if found
func (t *Tree) Find(value int) *TreeNode {
	panic("Tree.Find(int) not implemented")
	return nil
}

// Remove removes value if find and rebuilds tree
func (t *Tree) Remove(value int) {
	panic("Tree.Remove(value) not implemented")
}

// Insert insert value in tree
func (t *Tree) Insert(value int) {

	t.mutex.Lock()

	if t.head == nil {

		t.head = &TreeNode{value: value, left: nil, right: nil, mutex: t.mutex}
		t.mutex.Unlock()
	} else {

		t.mutex.Unlock()
		if t.head.value == value {
			return
		}

		t.head.insert(value)

	}

}

func (tn *TreeNode) insert(value int) {

	tn.mutex.Lock()

	if tn.value == value {
		tn.mutex.Unlock()
		return
	}

	if value <= tn.value {
		if tn.left == nil {
			tn.left = &TreeNode{value: value, left: nil, right: nil, mutex: tn.mutex}
			tn.mutex.Unlock()
		} else {
			tn.mutex.Unlock()
			tn.left.insert(value)
		}
	} else {
		if tn.right == nil {
			tn.right = &TreeNode{value: value, left: nil, right: nil, mutex: tn.mutex}
			tn.mutex.Unlock()
		} else {
			tn.mutex.Unlock()
			tn.right.insert(value)
		}
	}
}
