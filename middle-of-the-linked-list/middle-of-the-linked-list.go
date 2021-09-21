/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    count := 1
    
    curNode := head
    for curNode.Next != nil {
        curNode = curNode.Next
        count++
    }
    
    middle := middle(count)
    fmt.Println(middle)
    curNode = head
    for i := 1; i < middle; i++ {
        curNode = curNode.Next
    }
    
    return curNode
}

func middle(count int) int {
    if count % 2 == 0 {
        return count / 2 + 1
    }
    return (count + 1) / 2
}