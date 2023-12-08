func maxTaxiEarnings(n int, rides [][]int) int64 {
    sort.Slice(rides, func(x, y int)bool{
        if rides[x][1] == rides[y][1]{
            return rides[x][2] > rides[y][2]
        }
        return rides[x][1] < rides[y][1]
    })

    dp := make([]int, len(rides) + 1)
    for i := 0; i < len(rides); i++{
        j := sort.Search(i+1, func(k int)bool{
            return rides[k][1] > rides[i][0]
        })
        dp[i+1] = max(dp[i], dp[j] + rides[i][1] - rides[i][0] + rides[i][2])
    }
    return int64(dp[len(rides)])
}  