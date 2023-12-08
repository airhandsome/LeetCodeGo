func maxTaxiEarnings(n int, rides [][]int) int64 {
    dp := make([]int64, n + 1)
    rideMap := map[int][][]int{}
    for _, ride := range rides {
        rideMap[ride[1]] = append(rideMap[ride[1]], ride)
    }
    for i := 1; i <= n; i++ {
        dp[i] = dp[i - 1]
        for _, ride := range rideMap[i] {
            dp[i] = max(dp[i], dp[ride[0]] + int64(ride[1] - ride[0] + ride[2]))
        }
    }
    return dp[n]
}
