func minPathCost(grid [][]int, moveCost [][]int) int {
    m, n := len(grid), len(grid[0])
    fir := grid[0]
    for i := 1; i < m; i++{
        sec := make([]int, n)
        for j := 0; j < n; j ++{
            sec[j] = 1000000
            for k := 0; k < n; k++{
                sec[j] = min(sec[j], fir[k] + moveCost[grid[i-1][k]][j] + grid[i][j])
            }            
        }
        fir = sec
    }

    ans := 1000000
    for _, v := range fir{
        ans = min(ans, v)
    }
    return ans
}