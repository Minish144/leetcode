/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    var cur *ListNode = head
    
    for cur != nil {
        prev, cur, cur.Next = cur, cur.Next, prev
    }
    
    return prev
}