func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currentSum := nums[0]
    
    for _, num := range nums[1:len(nums)] {
        currentSum = max(currentSum + num, num)
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}