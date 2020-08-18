package bstlogger

import (
	"binarytree/pkg/logging"
	"binarytree/pkg/tree/binary"
	"testing"
)

func TestBstLogger_Insert(t *testing.T) {
	t.Run("must write value into inner bstree", func(t *testing.T) {
		testBst := &testBSTree{}
		testLogger := &testLogger{}
		bstlogged := NewBSTLogger(testBst, testLogger)

		expectedValue := 1337

		bstlogged.Insert(expectedValue)

		if testBst.value != expectedValue {
			t.Errorf("Insert(value) must write value into inner bst, expected %d , got %d", expectedValue, testBst.value)
		}

		if len(testLogger.values) == 0 {
			t.Error("Insert(value) must write logs data into inner logger, got 0 fields, expected not 0")
		}
	})
}

type testBSTree struct {
	value int
}

func (t *testBSTree) Insert(value int) {
	t.value = value
}

func (t *testBSTree) Remove(value int) {
	t.value = value
}

func (t *testBSTree) Find(value int) *binary.TreeNode {
	t.value = value
	return nil
}

type testLogger struct {
	values map[string]interface{}
	LoggerWithFields
}

func (tl *testLogger) ErrorWithFields(f logging.Fields) {
	tl.values = f
}

func (tl *testLogger) InfoWithFields(f logging.Fields) {
	tl.values = f
}

func TestBstLogger_Remove(t *testing.T) {
	t.Run("must write value into inner bstree", func(t *testing.T) {
		testBst := &testBSTree{}
		testLogger := &testLogger{}
		bstlogged := NewBSTLogger(testBst, testLogger)

		expectedValue := 1337

		bstlogged.Remove(expectedValue)

		if testBst.value != expectedValue {
			t.Errorf("Remove(value) must write value into inner bst, expected %d , got %d", expectedValue, testBst.value)
		}

		if len(testLogger.values) == 0 {
			t.Error("Remove(value) must write logs data into inner logger, got 0 fields, expected not 0")
		}
	})
}

func TestBstLogger_Find(t *testing.T) {
	t.Run("must write value into inner bstree", func(t *testing.T) {
		testBst := &testBSTree{}
		testLogger := &testLogger{}
		bstlogged := NewBSTLogger(testBst, testLogger)

		expectedValue := 1337

		bstlogged.Find(expectedValue)

		if testBst.value != expectedValue {
			t.Errorf("Find(value) must write value into inner bst, expected %d , got %d", expectedValue, testBst.value)
		}

		if len(testLogger.values) == 0 {
			t.Error("Find(value) must write logs data into inner logger, got 0 fields, expected not 0")
		}
	})
}
