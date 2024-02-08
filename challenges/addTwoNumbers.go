/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    tmp := result
    for l1 != nil || l2 != nil {
        if l1 != nil {
            tmp.Val += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            tmp.Val += l2.Val
            l2 = l2.Next
        }

        if tmp.Val > 9 {
            tmp.Val -= 10
            tmp.Next = &ListNode{Val: 1}
        } else if l1 != nil || l2 != nil {
            tmp.Next = &ListNode{}
        }
        tmp = tmp.Next
    }
    return result
}
