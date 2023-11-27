func sumSubarrayMins(arr []int) int {
    n := len(arr)
    left, right := make([]int, n), make([]int, n)

    left[0] = -1
    right[n-1] = n
    for i := 1; i < n; i++{
        index := i - 1    
        for index != -1 && arr[index] >= arr[i]{
            index = left[index]
        }
        left[i] = index

        index2 := n - i 
        for index2 != n && arr[index2] > arr[n-i-1]{
            index2 = right[index2]
        }
        right[n-i-1] = index2
    }

    ans := 0
    mod := int(1e9 + 7)
    for i := 0; i < n; i++{
        // fmt.Println(i-left[i], right[i]-i, arr[i])
        ans += (i - left[i]) * (right[i] - i) * arr[i]
        ans %= mod
    }
    return ans
}