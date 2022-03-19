/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    counter := make(map[*ListNode]int)
    cur := head
    for cur != nil {
        if _, ok := counter[cur]; ok {
            return true
        } else {
            counter[cur]++
        }
        cur = cur.Next
    }
    
    return false
}