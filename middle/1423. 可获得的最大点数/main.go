func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    sum := 0

    for _, v := range cardPoints[:n-k]{
        sum += v
    }    
    
    minimun := sum
    for i := n-k; i < n; i++{
        sum = sum - cardPoints[i + k - n] + cardPoints[i]
        minimun = min(minimun, sum)
    }

    ans := 0
    for _, v := range cardPoints{
        ans += v
    }
    return ans - minimun

}