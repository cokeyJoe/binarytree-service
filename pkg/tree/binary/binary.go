// Package binary contains binary tree implementation
package binary

import (
	"encoding/json"
	"sync"
)

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

func (tn *TreeNode) MarshalJSON() ([]byte, error) {
	if tn == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(tn.value)
}

// New returns new binary tree with nodes with values
func New(values ...int) *Tree {

	t := new()

	for _, v := range values {
		t.Insert(v)
	}

	return t
}

// new returns new empty binary tree
func new() *Tree {
	mutex := &sync.Mutex{}
	return &Tree{mutex: mutex}
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

	if t.head == nil {
		return
	}

	t.head.remove(value)
}

func (tn *TreeNode) remove(value int) *TreeNode {

	if tn == nil {
		return nil
	}

	if value < tn.value {
		tn.left = tn.left.remove(value)
		return tn
	}
	if value > tn.value {
		tn.right = tn.right.remove(value)
		return tn
	}

	if tn.left == nil && tn.right == nil {
		tn = nil
		return nil
	}

	if tn.left == nil {
		tn = tn.right
		return tn
	}
	if tn.right == nil {
		tn = tn.left
		return tn
	}

	minValueNode := tn.right
	for {

		if minValueNode != nil && minValueNode.left != nil {
			minValueNode = minValueNode.left
		} else {
			break
		}
	}

	tn.value = minValueNode.value
	tn.right = tn.right.remove(tn.value)
	return tn
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

func (tn *TreeNode) Value() int {
	if tn == nil {
		return 0
	}

	return tn.value
}
