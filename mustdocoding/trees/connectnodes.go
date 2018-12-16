package mdc

import "fmt"

//SiblingTree ... represents tree with siblings
type SiblingTree struct {
	Data    int
	Left    *SiblingTree
	Right   *SiblingTree
	Sibling *SiblingTree
}

type queueData struct {
	level int
	sb    *SiblingTree
}

func (sb *SiblingTree) ConnectNodesAtSameLevel() {
	queue := []*queueData{}
	var tree *queueData
	queue = append(queue, &queueData{level: 0, sb: sb})
	currentLevel := -1
	var previousSibling *SiblingTree
	for len(queue) > 0 {
		tree = queue[0]
		if currentLevel < tree.level {
			currentLevel = tree.level
		} else {
			previousSibling.Sibling = tree.sb
		}

		previousSibling = tree.sb

		queue = queue[1:]
		if tree.sb.Left != nil {
			queue = append(queue, &queueData{level: tree.level + 1, sb: tree.sb.Left})
		}

		if tree.sb.Right != nil {
			queue = append(queue, &queueData{level: tree.level + 1, sb: tree.sb.Right})
		}
	}
}

func (sb *SiblingTree) LevelOrderTraverse() {
	for levelStart := sb; levelStart != nil; levelStart = levelStart.Left {
		for sibling := levelStart; sibling != nil; sibling = sibling.Sibling {
			fmt.Printf("%d ", sibling.Data)
		}
	}
	fmt.Println()
}
