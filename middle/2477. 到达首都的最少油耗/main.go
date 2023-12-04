func minimumFuelCost(roads [][]int, seats int) int64 {
    link := make([][]int, len(roads)+1)
    for _, v := range roads{
        if len(link[v[0]]) == 0{
            link[v[0]] = []int{}
        }
        if len(link[v[1]]) == 0{
            link[v[1]] = []int{}
        }        
        link[v[0]] = append(link[v[0]], v[1])
        link[v[1]] = append(link[v[1]], v[0])
    }

    res := 0
    var dfs func(cur, father int) int
    dfs = func(cur, father int)int{
        peopleSum := 1
        for _, next := range link[cur]{
            if next != father{
                peopleCnt := dfs(next, cur)
                peopleSum += peopleCnt
                res += (peopleCnt + seats - 1) /seats
            }
        }
            
        return peopleSum
    }
    dfs(0, -1)
    return int64(res)
}