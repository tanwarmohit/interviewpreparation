package mdc

import "testing"

func TestCheckBst(t *testing.T) {
	//n := 4
	type sample struct {
		parent, child, direction int
	}

	//10 20 L 10 30 R 20 40 L 20 60 R
	testData := []sample{
		{10, 20, 1},
		{20, 30, 1},
		{30, 40, 1},
		{40, 60, 1},
	}

	inputMap := make(map[int]*BinaryTree)
	r := 0
	for _, td := range testData {
		p, ok := inputMap[td.parent]
		if r == 0 {
			r = td.parent
		}
		if !ok {
			inputMap[td.parent] = &BinaryTree{Data: td.parent, Left: nil, Right: nil}
			p = inputMap[td.parent]
		}
		inputMap[td.child] = &BinaryTree{Data: td.child, Left: nil, Right: nil}
		if td.direction == 0 {
			p.Left = inputMap[td.child]
		} else {
			p.Right = inputMap[td.child]
		}
	}

	inputMap[r].inOrderTraversal()
	if inputMap[r].CheckBst() == false {
		t.Fail()
	}
}
