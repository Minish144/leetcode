func missingNumber(nums []int) int {
    numberMap := make(map[int]int)
    for _, num := range nums {
        numberMap[num]++
    }
    
    for i := 0; i <= len(nums); i++ {
        fmt.Println("i:", i)
        if _, ok := numberMap[i]; !ok {
            return i
        }
    }
    return 0
}