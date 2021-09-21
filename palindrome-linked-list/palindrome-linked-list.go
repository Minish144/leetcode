/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {   
    node := head
    
    vals := []int{node.Val}
    for node.Next != nil {
        vals = append(vals, node.Next.Val)
        node = node.Next
    }
    
    count := len(vals)
    if count <= 1 {
        return true
    }
    
    for i := 0; i < count / 2; i++ {
        if vals[i] != vals[len(vals) - 1 - i] {
            return false
        }
    }

    return true
}