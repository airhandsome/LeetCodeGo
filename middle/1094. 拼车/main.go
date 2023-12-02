func carPooling(trips [][]int, capacity int) bool {
    station := make([]int, 1001)
    for _, v := range trips{
        station[v[1]] += v[0]
        station[v[2]] -= v[0]
    }

    for _, v := range station{
        capacity -= v
        if capacity < 0{
            return false
        }
    }
    return true
}