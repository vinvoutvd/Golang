//https://leetcode.com/problems/path-sum/?envType=study-plan-v2&envId=top-interview-150
//Definition for a binary tree node.
type TreeNode struct {
    Val int
	Left *TreeNode
	Right *TreeNode
}

// решение с помощью рекурсии и добавления дополнительного параметра total для хранения промежуточных значений
func dfs (node *TreeNode, total int, target int) (bool){
		if node == nil {
			return false
		}
		total+=node.Val
		if node.Left==nil && node.Right==nil{
			return total == target
		}
	return dfs(node.Left, total, target) || dfs(node.Right, total, target)
	}
 
func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)} 
