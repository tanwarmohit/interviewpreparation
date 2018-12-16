package mdc

import (
	"fmt"
)

type BinaryTree struct {
	Data  int
	Left  *BinaryTree
	Right *BinaryTree
}

func (bt *BinaryTree) checkBstInternal(leftBound int, rightBound int) bool {
	// base case: If node is null, then its a BST
	if bt == nil {
		return true
	}

	return bt.Data > leftBound &&
		bt.Data < rightBound &&
		bt.Left.checkBstInternal(leftBound, bt.Data) &&
		bt.Right.checkBstInternal(bt.Data, rightBound)
}

//CheckBst ... check whether bt is bst or not
func (bt *BinaryTree) CheckBst() bool {
	return bt.checkBstInternal(-1001, 1001)
}

func (bt *BinaryTree) inOrderTraversal() {
	if bt == nil {
		return
	}
	bt.Left.inOrderTraversal()
	fmt.Printf("%d\n", bt.Data)
	bt.Right.inOrderTraversal()
}
