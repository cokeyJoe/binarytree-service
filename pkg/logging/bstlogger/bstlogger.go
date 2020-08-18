package bstlogger

import (
	"binarytree/pkg/logging"
	"binarytree/pkg/tree/binary"
	"time"
)

type BstLogger struct {
	innerBst BinarySearchTree
	logger   LoggerWithFields
}

type LoggerWithFields interface {
	ErrorWithFields(logging.Fields)
	InfoWithFields(logging.Fields)
}

type BinarySearchTree interface {
	Find(int) *binary.TreeNode
	Remove(int)
	Insert(int)
}

// NewBSTLogger returns BinarySearchTree wrapped into logger
func NewBSTLogger(bst BinarySearchTree, logger LoggerWithFields) *BstLogger {
	return &BstLogger{
		innerBst: bst,
		logger:   logger,
	}
}

func (l *BstLogger) Find(value int) *binary.TreeNode {

	start := time.Now()

	node := l.innerBst.Find(value)
	ms := time.Since(start).Milliseconds()

	l.logger.InfoWithFields(logging.Fields{
		"input_value": value,
		"elapsed_ms":  ms,
		"result":      node,
		"action":      "Find",
	})

	return node
}

func (l *BstLogger) Remove(value int) {
	start := time.Now()

	l.innerBst.Remove(value)

	ms := time.Since(start).Milliseconds()
	l.logger.InfoWithFields(logging.Fields{
		"input_value": value,
		"elapsed_ms":  ms,
		"action":      "Remove",
	})
}

func (l *BstLogger) Insert(value int) {
	start := time.Now()

	l.innerBst.Insert(value)

	ms := time.Since(start).Milliseconds()

	l.logger.InfoWithFields(logging.Fields{
		"input_value": value,
		"elapsed_ms":  ms,
		"action":      "Insert",
	})
}
