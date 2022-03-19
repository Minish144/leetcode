func climbStairs(n int) int {
    var a, b int = 1, 1
    for i := 0; i < n; i++ {
        a, b = b, a+b
    }
    return a
}
