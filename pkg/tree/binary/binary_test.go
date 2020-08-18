// Package binary contains binary tree implementation
package binary

import (
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
		tree := TreeNode{value: 100, left: nil, right: nil}
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
		tree := TreeNode{value: 100, left: nil, right: nil}
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
		tree := TreeNode{value: 100, left: &TreeNode{value: 50}}
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
		tree := TreeNode{value: 100, left: nil, right: nil}
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
		tree := TreeNode{value: 100, right: &TreeNode{value: 150}}
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

func TestNewFromInts(t *testing.T) {
	t.Run("must return valid tree", func(t *testing.T) {

		expectedValue_59 := 59
		tree := New(expectedValue_59, 12, 3, 99, 29, 22, 491, 213, 5, 6, 223, 12)

		if tree.head == nil {
			t.Error("NewFromInts must return not nil tree.head, got nil")
			t.FailNow()
		}

		if tree.head.value != expectedValue_59 {
			t.Errorf("tree.head.value expected %d, got %d", expectedValue_59, tree.head.value)
		}
	})
}

func TestTree_Find(t *testing.T) {
	t.Run("Tree's head is nil, must return nil", func(t *testing.T) {
		tree := New()

		val := tree.Find(123)

		if val != nil {
			t.Error("Tree.Search(int) must return nil, got not nil")
		}
	})
}

func TestTreeNode_find(t *testing.T) {
	t.Run("Must find left tree", func(t *testing.T) {

		tree := New(50, 10, 40)

		expectedValue := 10
		node := tree.Find(expectedValue)

		if node == nil {
			t.Error("expected not nil, got nil", node)
			t.FailNow()
		}

		if node.value != expectedValue {
			t.Errorf("node value expected %d, got %d", node.value, expectedValue)
		}

	})

	t.Run("Tree has no left node, must return nil", func(t *testing.T) {
		tree := New(50, 60, 66)

		node := tree.Find(10)

		if node != nil {
			t.Error("expected nil, got not nil")
			t.FailNow()
		}

	})

	t.Run("Must find right tree", func(t *testing.T) {

		tree := New(50, 10, 40, 60)

		expectedValue := 60
		node := tree.Find(expectedValue)

		if node == nil {
			t.Error("expected not nil, got nil", node)
			t.FailNow()
		}

		if node.value != expectedValue {
			t.Errorf("node value expected %d, got %d", node.value, expectedValue)
		}

	})

	t.Run("Tree has no right node, must return nil", func(t *testing.T) {
		tree := New(50, 15, 10)

		node := tree.Find(60)

		if node != nil {
			t.Error("expected nil, got not nil")
			t.FailNow()
		}

	})
}

func TestTreeNode_remove(t *testing.T) {
	t.Run("Head is nil, must return", func(t *testing.T) {
		tree := New()

		defer func() {
			if recover() != nil {
				t.Error("must not panic")
			}
		}()
		tree.Remove(1)

	})

	t.Run("Head has value, must not affect", func(t *testing.T) {
		tree := New(1)

		tree.Remove(2)

		defer func() {
			if recover() != nil {
				t.Error("must not panic")
			}
		}()
	})

	t.Run("must remove left node", func(t *testing.T) {
		tree := New(5, 4, 6)

		tree.Remove(4)

		defer func() {
			if recover() != nil {
				t.Error("must not panic")
			}
		}()

		if tree.head == nil {
			t.Error("Expected tree.head will not be removed")
			t.FailNow()
		}

		if tree.head.left != nil {
			t.Error("expected tree.head.left will be removed")
		}
	})

	t.Run("must replace first left with last rights left node", func(t *testing.T) {
		tree := New(100, 75, 60, 85, 81)

		_ = tree

		tree.Remove(75)

		if tree.head.left.value != 81 {
			t.Error("expected node replace")
		}
	})
}

func TestTreeNode_Value(t *testing.T) {
	type fields struct {
		value int
		left  *TreeNode
		right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "is 5, must return 5", fields: fields{value: 5}, want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := &TreeNode{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tn.Value(); got != tt.want {
				t.Errorf("TreeNode.Value() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("is nil, must return 0", func(t *testing.T) {
		var tree *TreeNode

		val := tree.Value()

		defer func() {
			if recover() != nil {
				t.Error("must not painc")
			}
		}()

		if val != 0 {
			t.Errorf("must return zero value, got %d", val)
		}
	})
}
