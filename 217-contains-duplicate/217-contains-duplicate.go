func containsDuplicate(nums []int) bool {
    counterMap := make(map[int]int)
    for _, num := range nums {
        counterMap[num]++
        if containCount := counterMap[num]; containCount > 1 {
            return true
        }
    }
    return false
}