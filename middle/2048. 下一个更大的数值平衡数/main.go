func nextBeautifulNumber(n int) int {
    check := func(x int)bool{
        count := make([]int, 10)
        for x > 0{
            count[x % 10] ++
            x /= 10
        }
        for i := 0; i < 10; i++{
            if count[i] > 0 && count[i] != i{
                return false
            }
        }
        return true
    }
    for i := n+1; i <=1224444; i++{
        if check(i){
            return i
        }
    }
    return -1
}