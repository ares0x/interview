package tree

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val: val,
	}
}

type BinaryTree struct {
	root *TreeNode
}

// 前序 根 -> 左 -> 右
func (t *BinaryTree) Preorder() []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.val)
		preorder(node.left)
		preorder(node.right)
	}
	preorder(t.root)
	return res
}

// 中序 左 -> 根 -> 右
func (t *BinaryTree) InOrder() []int {
	res := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.left)
		res = append(res, node.val)
		inorder(node.right)
	}
	inorder(t.root)
	return res
}

// 后序 左 -> 右 -> 根
func (t *BinaryTree) PostOrder() []int {
	res := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.left)
		postorder(node.right)
		res = append(res, node.val)
	}
	postorder(t.root)
	return res
}

func (t *BinaryTree) Insert(val int) {
	n := NewTreeNode(val)
	if t.root == nil {
		t.root = n
	} else {
		cur := t.root
		for {
			if val < cur.val {
				if cur.left == nil {
					cur.left = n
					break
				}
				cur = cur.left
			} else {
				if cur.right == nil {
					cur.right = n
					break
				}
				cur = cur.right
			}
		}
	}
}

func (t *BinaryTree) Search(val int) *TreeNode {
	cur := t.root
	for cur != nil {
		if cur.val == val {
			return cur
		} else if val < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}
