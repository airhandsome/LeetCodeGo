func minDeletion(nums []int) int {
    n := len(nums)
    index := 0
    ans := 0
    for i, v := range nums[:n-1]{
        if index % 2 == 0 && v == nums[i+1]{
            ans ++
        }else{
            index ++
        }
    }
    if index% 2 == 0{
        ans ++
    }
    return ans
}