/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pseudoPalindromicPaths (root *TreeNode) int {
    counter := make([]int, 10)

    var dfs  func(node *TreeNode, counter []int) int

    dfs = func(node *TreeNode, counter []int) int{
        if node == nil{
            return 0
        }
        counter[node.Val]++
        ans := 0
        if node.Left == nil && node.Right == nil{
            odd := 0
            for _, v := range counter{
                if v % 2 == 1{
                    odd ++
                }
            }
            if odd <= 1{
                ans ++
            }
        }else{
            ans = dfs(node.Left, counter) + dfs(node.Right, counter)
        }
        counter[node.Val] --
        return ans
    }

    return dfs(root, counter)
}