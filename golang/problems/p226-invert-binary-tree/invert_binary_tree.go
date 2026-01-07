package solution

import (
	. "leetcode/golang/base"
)

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var dfs = func(*TreeNode){}

    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }

        root.Left, root.Right = root.Right, root.Left
        dfs(root.Left)
        dfs(root.Right)      
    }

    dfs(root)
    return root
}