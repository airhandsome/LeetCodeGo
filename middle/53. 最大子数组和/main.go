func maxSubArray(nums []int) int {
    ans := nums[0]
    sum := 0
    for _, v := range nums{
        sum += v
        ans = max(ans, sum)
        if sum < 0{
            sum = 0
        }
    }
    return ans
}