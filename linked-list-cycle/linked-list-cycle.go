/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Floyd's tortoise and hare
// https://en.wikipedia.org/wiki/Cycle_detection

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    fast := head
    slow := head
    
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        
        if fast == slow {
            return true
        }
    }
    
    return false
}