func firstCompleteIndex(arr []int, mat [][]int) int {
    m, n := len(mat), len(mat[0])
    position := make([][2]int, m*n+1)

    for i, row := range mat{
        for j, v := range row{
            position[v] = [2]int{i,j}
        }
    }

    row, col := make([]int, m), make([]int, n)

    for i, v := range arr{
        x, y := position[v][0], position[v][1]
        row[x]++
        col[y]++
        if row[x] == n || col[y] == m{
            return i
        }

    }
    return -1
}