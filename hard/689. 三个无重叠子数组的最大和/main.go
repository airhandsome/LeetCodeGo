func maxSumOfThreeSubarrays(nums []int, k int) []int {
    ans := []int{}
    sum1, sum2, sum3 := 0, 0, 0
    maxSum1, maxSum2, maxSum3 := 0, 0, 0
    id1, id21, id22 := 0, 0, 0
    for i := 2 * k; i < len(nums); i++{
        sum1 += nums[i - 2 * k]
        sum2 += nums[i - k]
        sum3 += nums[i]
        if i >= 3 * k - 1{
            if sum1 > maxSum1{
                maxSum1 = sum1
                id1 = i - k * 3 + 1
            }
            if maxSum1 + sum2 > maxSum2{
                maxSum2 = sum2 + maxSum1
                id21 = id1
                id22 = i - k * 2 + 1                
            }
            if maxSum2 + sum3 > maxSum3{
                maxSum3 = sum3 + maxSum2
                ans = []int{id21, id22, i - k + 1}
            }            
            sum1 -= nums[i - k * 3 + 1]
            sum2 -= nums[i - k * 2 + 1]
            sum3 -= nums[i - k + 1]
        }
    }
    return ans
}