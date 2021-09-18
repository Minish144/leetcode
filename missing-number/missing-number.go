func missingNumber(nums []int) int {
    n := len(nums)
    sSum := seriesSum(0, n + 1)
    sum := sum(nums)
    fmt.Println(n + 1, sSum, sum)
    return sSum - sum   
}

func sum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func seriesSum(a_1 int, n int) int {
    return (n - 1) * n / 2
}

func max(nums []int, leng int) int {
    sort.Ints(nums)
    return nums[leng - 1]
}