func containsDuplicate(nums []int) bool {
    unique := []int{}
    for _, num := range nums {
        if contains(unique, num) {
            return true
        }
        unique = append(unique, num)
    }
    return false
}

func contains(arr []int, value int) bool {
    for _, elem := range arr {
        if elem == value {
            return true
        }
    }
    return false
}
