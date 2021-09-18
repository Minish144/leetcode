func singleNumber(nums []int) int {
    hashmap := make(map[int]int)
    n := len(nums)
    
    for i := 0; i < n; i++ {
        hashmap[nums[i]]++
    }
    
    for i := 0; i < n; i++ {
        fmt.Println(i, nums[i], hashmap[nums[i]])
        if val, _ := hashmap[nums[i]]; val == 1 {
            return nums[i]
        }
    }
    return 0
}