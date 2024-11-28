package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	p1 := head
	p2 := head.Next

	ans := p2

	for p1 != nil && p2 != nil {
		swapNode(prev, p1, p2)
		prev = p1
		p1 = p1.Next

		if p1 != nil {
			p2 = p1.Next
		}
	}

	return ans
}

func swapNode(prev *ListNode, p1 *ListNode, p2 *ListNode) {
	if prev == nil {
		p1.Next = p2.Next
		p2.Next = p1
		return
	}

	prev.Next = p1.Next
	p1.Next = p2.Next
	p2.Next = p1
}

func main() {
	p4 := &ListNode{
		Val: 4,
	}

	p3 := &ListNode{
		Val:  3,
		Next: p4,
	}

	p2 := &ListNode{
		Val:  2,
		Next: p3,
	}

	head := &ListNode{
		Val:  1,
		Next: p2,
	}

	travel := swapPairs(head)

	for travel != nil {
		fmt.Print(fmt.Sprintf("%v%s", travel.Val, " -> "))
		travel = travel.Next
	}
}
