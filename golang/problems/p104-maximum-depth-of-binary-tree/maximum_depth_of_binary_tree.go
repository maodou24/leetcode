package solution

import . "leetcode/golang/base"

func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode, int)

	var ans int
	dfs = func(node *TreeNode, n int) {
		if node == nil {
			return
		}

		ans = max(ans, n)

		if node.Left != nil {
			dfs(node.Left, n+1)
		}
		if node.Right != nil {
			dfs(node.Right, n+1)
		}
	}

	dfs(root, 1)
	return ans
}
