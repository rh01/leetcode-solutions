package add_two_nums

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// [2[4[3]]]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// var m []int
	// for ; l1.Next != nil; l1 = l1.Next {
	// 	m = append(m, l1.Val)
	// }

	// reuturn

	var carry int
	dummyHead := new(ListNode) // head of result list
	curr := dummyHead
	var p, q = l1, l2

	for p != nil || q != nil {
		var x, y, sum int
		if p == nil {
			x = 0
		} else {
			x = p.Val
			p = p.Next
		}
		if q == nil {
			y = 0
		} else {
			y = q.Val
			q = q.Next
		}

		sum = x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next

}
