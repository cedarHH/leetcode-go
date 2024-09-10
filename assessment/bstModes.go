package assessment

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newBst(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func findMode(root *TreeNode) []int {
	var base, count, maxCount int
	var result []int

	update := func(val int) {
		if val != base {
			base = val
			count = 0
		}
		count++
		if count == maxCount {
			result = append(result, val)
		} else if count > maxCount {
			maxCount = count
			result = []int{val}
		}
	}
	curr := root
	for curr != nil {
		if curr.Left != nil {
			prev := curr.Left
			for ; prev.Right != nil; prev = prev.Right {
				if prev.Right == curr {
					break
				}
			}
			if prev.Right == nil {
				prev.Right = curr
				curr = curr.Left
			} else {
				prev.Right = nil
				update(curr.Val)
				curr = curr.Right
			}
		} else {
			update(curr.Val)
			curr = curr.Right
		}
	}
	return result
}

func FindMode() {
	arr := []int{1, -1, 2, 2, 3, 3}
	bst := newBst(arr[0])
	curr := bst
	index := 1

	var insert func(node *TreeNode)
	insert = func(node *TreeNode) {
		if index >= len(arr) {
			return
		} else if arr[index] != -1 && arr[index-1] != -1 {
			node.Left = newBst(arr[index])
			index++
			insert(node.Left)
		} else if arr[index-1] == -1 {
			node.Right = newBst(arr[index])
			index++
			insert(node.Right)
		} else if arr[index] == -1 {
			index++
			insert(node)
		}
	}

	insert(curr)

	fmt.Printf("%v\n", findMode(bst))
}
