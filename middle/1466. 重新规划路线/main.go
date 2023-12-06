func minReorder(n int, connections [][]int) int {
    var dfs func(x, parent int, c [][][]int)int

    dfs = func(x, parent int, c [][][]int)int{
        ans := 0
        for _, edge := range c[x]{
            if edge[0] == parent{
                continue
            }
            ans += edge[1] + dfs(edge[0], x, c)
        }
        return ans
    }

    c := make([][][]int, n)
    for _, edge := range connections{
        c[edge[0]] = append(c[edge[0]], []int{edge[1], 1})
        c[edge[1]] = append(c[edge[1]], []int{edge[0], 0})
    }
    return dfs(0, -1, c)
}