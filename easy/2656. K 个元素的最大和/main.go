func maximizeSum(nums []int, k int) int {
    maxVal := 0
    for _, v := range nums{
        maxVal = max(maxVal, v)
    }
    return maxVal * k + (k -1) * k / 2
} 