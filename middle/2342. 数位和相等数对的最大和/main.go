func maximumSum(nums []int) int {
    fir, sec := map[int]int{}, map[int]int{}
    ans := -1
    for _, v := range nums{
        val := v
        sum := 0
        for val > 0{
            sum += val % 10
            val /= 10
        }
        if v > fir[sum]{
            sec[sum], fir[sum] = fir[sum], v
        }else if v > sec[sum]{
            sec[sum] = v
        }
        if sec[sum] > 0{
            ans = max(ans, fir[sum] + sec[sum])
        }
    }
    return ans

}