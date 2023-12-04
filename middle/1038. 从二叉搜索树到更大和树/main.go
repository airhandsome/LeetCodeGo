func bstToGst(root *TreeNode) *TreeNode {
    maxTree(root, 0)
    return root

}

func maxTree(root *TreeNode, parentVal int) int{
    if root == nil{
        return 0
    }
    rightMax := maxTree(root.Right, parentVal)
    root.Val += max(rightMax, parentVal)
    leftMax := maxTree(root.Left, root.Val)
    
    return max(leftMax, root.Val)
}