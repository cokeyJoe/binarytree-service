// Package binary contains binary tree implementation
package binary

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		wantNil bool
	}{
		{
			name:    "must return not nil",
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); got == nil {
				t.Errorf("New() = %v, want %v", got, tt.wantNil)
			}
		})
	}

	t.Run("New Tree must not has values", func(t *testing.T) {
		tree := New()

		if tree.head != nil {
			t.Errorf("Expected empty tree, got %v", tree.head)
		}

	})
}

func TestTree_Insert(t *testing.T) {

	t.Run("New() Tree.Insert() must insert value into head node", func(t *testing.T) {
		tree := New()
		expectedInt := 1337
		tree.Insert(expectedInt)

		if tree.head == nil {
			t.Errorf("Tree.Insert(int) must insert TreeNode with value %v, got nil", expectedInt)
			t.FailNow()
		}

		if tree.head.value != expectedInt {
			t.Errorf("Tree.Insert(int) must insert TreeNode with value %d, got %d", expectedInt, tree.head.value)
		}
	})

	t.Run("New() Tree.Insert() must be idempotent", func(t *testing.T) {
		tree := New()
		expectedInt := 1337
		tree.Insert(expectedInt)

		tree.Insert(expectedInt)

		if tree.head == nil {
			t.Errorf("Tree.Insert(int) must insert TreeNode with value %v, got nil", expectedInt)
			t.FailNow()
		}

		if tree.head.value != expectedInt {
			t.Errorf("Tree.Insert(int) must insert TreeNode with value %d, got %d", expectedInt, tree.head.value)
		}

		if tree.head.left != nil {
			t.Error("left node must be nil")
		}
		if tree.head.right != nil {
			t.Error("right node must be nil")
		}
	})

	t.Run("Tree.Insert() must insert value into left node", func(t *testing.T) {
		tree := New()

		tree.Insert(200)

		expectedValue := 150
		tree.Insert(expectedValue)

		if tree.head == nil {
			t.Errorf("Tree.Insert(int) must insert TreeNode , got nil")
			t.FailNow()
		}

		if tree.head.left == nil {
			t.Error("Tree.Insert(int) must insert into head left node, got nil")
			t.FailNow()
		}

		if tree.head.left.value != expectedValue {
			t.Errorf("Tree.Insert(int) must insert value into head's left node, expected %d, got %d", expectedValue, tree.head.left.value)
		}
	})

}

func TestTreeNode_insert(t *testing.T) {
	t.Run("TreeNode.insert(int) must insert into left node", func(t *testing.T) {
		tree := TreeNode{value: 100, left: nil, right: nil, mutex: &sync.Mutex{}}
		value := 40
		tree.insert(value)

		if tree.left == nil {
			t.Errorf("got nil, expected left node not nil")
			t.FailNow()
		}

		if tree.left.value != value {
			t.Errorf("got %d, expected %d", tree.left.value, value)
		}
	})

	t.Run("TreeNode.insert(int) must be idempotent", func(t *testing.T) {
		tree := TreeNode{value: 100, left: nil, right: nil, mutex: &sync.Mutex{}}
		value := 100
		tree.insert(value)

		if tree.left != nil {
			t.Error("left node must be nil")
		}

		if tree.right != nil {
			t.Error("right node must be nil")
		}

	})

	t.Run("TreeNode.insert(int) must insert value into left node's left node", func(t *testing.T) {
		tree := TreeNode{value: 100, left: &TreeNode{value: 50, mutex: &sync.Mutex{}}, mutex: &sync.Mutex{}}
		expectedValue := 10

		tree.insert(expectedValue)

		if tree.left.left == nil {
			t.Error("expected not nil, got nil")
			t.FailNow()
		}

		if tree.left.left.value != expectedValue {
			t.Errorf("expected %d, got %d", expectedValue, tree.left.left.value)
		}
	})

	t.Run("TreeNode.insert(int) must insert into right node", func(t *testing.T) {
		tree := TreeNode{value: 100, left: nil, right: nil, mutex: &sync.Mutex{}}
		value := 140
		tree.insert(value)

		if tree.right == nil {
			t.Errorf("got nil, expected left node not nil")
			t.FailNow()
		}

		if tree.right.value != value {
			t.Errorf("got %d, expected %d", tree.right.value, value)
		}
	})

	t.Run("TreeNode.insert(int) must insert value into right node's right node", func(t *testing.T) {
		tree := TreeNode{value: 100, right: &TreeNode{value: 150, mutex: &sync.Mutex{}}, mutex: &sync.Mutex{}}
		expectedValue := 160

		tree.insert(expectedValue)

		if tree.right.right == nil {
			t.Error("expected not nil, got nil")
			t.FailNow()
		}

		if tree.right.right.value != expectedValue {
			t.Errorf("expected %d, got %d", expectedValue, tree.right.right.value)
		}
	})

}
