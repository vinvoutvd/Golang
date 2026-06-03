// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }

// Решение с помощью рекурсии
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
    leftHeight := maxDepth(root.Left)
    rightHeight := maxDepth(root.Right)
    if leftHeight >= rightHeight{
		  return leftHeight+1
	}
	return rightHeight+1
}
