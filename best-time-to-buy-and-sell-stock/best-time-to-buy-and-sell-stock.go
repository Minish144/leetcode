func maxProfit(prices []int) int {
    maxProfit := 0
    currentMin := prices[0]
    
    for _, price := range prices {
        currentMin = getMin(currentMin, price)
        maxProfit = getMax(maxProfit, price - currentMin)
    }
    
    return maxProfit
}

func getMin(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func getMax(a int, b int) int {
    if a > b {
        return a
    }
    return b
}