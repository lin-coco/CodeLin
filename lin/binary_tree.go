package lin

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	KV    struct {
		key int
		val string
	}
}

func NewBinarySearchTree(slice []int) *TreeNode {
	root := &TreeNode{}

	return root
}
