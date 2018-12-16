package main

import (
	"fmt"
	"interviewpreparation/mustdocoding/trees"
	"os"
)

func main() {
	if f, err := os.Open("input.txt"); err != nil {
		fmt.Printf("err while opening file: %v\n", err)
	} else {
		inputMap := make(map[int]*mdc.SiblingTree)
		defer f.Close()
		var T, N, p, c, r int
		direction := ""
		fmt.Fscanf(f, "%d", &T)
		fmt.Printf("Tst cases: %d\n", T)
		for i := 0; i < T; i++ {
			fmt.Fscanf(f, "%d", &N)
			fmt.Printf("n: %d\n", N)
			for n := 0; n < N; n++ {
				fmt.Fscanf(f, "%d", &p)
				if r == 0 {
					r = p
				}
				parent, ok := inputMap[p]
				if !ok {
					inputMap[p] = &mdc.SiblingTree{Data: p, Left: nil, Right: nil, Sibling: nil}
					parent = inputMap[p]
				}
				fmt.Fscanf(f, "%d", &c)
				inputMap[c] = &mdc.SiblingTree{Data: c, Left: nil, Right: nil, Sibling: nil}
				fmt.Fscanf(f, "%s", &direction)
				if direction[0] == 'L' {
					parent.Left = inputMap[c]
				} else {
					parent.Right = inputMap[c]
				}
			}
			inputMap[r].ConnectNodesAtSameLevel()
			inputMap[r].LevelOrderTraverse()
			r = 0
		}
	}
}
