package api

import (
	"binarytree/pkg/tree/binary"
)

type Finder interface {
	Find(int) *binary.TreeNode
}

type Inserter interface {
	Insert(int)
}

type Remover interface {
	Remove(int)
}

type BinarySearchTree interface {
	Finder
	Inserter
	Remover
}
