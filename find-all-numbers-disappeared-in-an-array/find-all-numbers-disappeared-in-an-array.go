func findDisappearedNumbers(nums []int) []int {
    hashmap := make(map[int]int)
    n := len(nums)
    
    res := []int{}
    
    for i := 0; i < n; i++ {
        hashmap[nums[i]]++
    }
    
    for i := 1; i <= n; i++ {
        if _, ok := hashmap[i]; !ok {
            res = append(res, i)
        }
    }
    
    return res
}