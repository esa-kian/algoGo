/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []

*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeListsDivideAndConquer(lists, 0, len(lists)-1)
}

func mergeListsDivideAndConquer(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	mid := start + (end-start)/2
	left := mergeListsDivideAndConquer(lists, start, mid)
	right := mergeListsDivideAndConquer(lists, mid+1, end)

	return mergeLists(left, right)
}

func mergeLists(listA, listB *ListNode) *ListNode {
	res := &ListNode{}
	dummyHead := res
	for listA != nil || listB != nil {
		if listA == nil {
			res.Next = listB
			break
		}
		if listB == nil {
			res.Next = listA
			break
		}

		if listA.Val < listB.Val {
			res.Next = listA
			res = res.Next
			listA = listA.Next
		} else {
			res.Next = listB
			res = res.Next
			listB = listB.Next
		}
	}
	return dummyHead.Next
}
