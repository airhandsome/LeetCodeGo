func climbStairs(n int) int {
    if n == 1{
        return 1
    }
    a, b := 0, 1
    for i := 0; i <= n; i++{
        a, b = b, a + b 
    }
    return a

}