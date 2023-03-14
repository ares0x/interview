package tree

// binary search tree(bst)：二叉搜索树
// 左子树小于根节点，右子树大于根节点
func (t *BinaryTree) Insert(val int) {
	node := NewTreeNode(val)
	if t.root == nil {
		t.root = node
		return
	}
	cur := t.root
	for {
		if val < cur.val {
			if cur.left == nil {
				cur.left = node
				break
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = node
				break
			}
			cur = cur.right
		}
	}
}

func (t *BinaryTree) Search(val int) *TreeNode {
	if t.root == nil {
		return nil
	}
	cur := t.root
	for {
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

// 删除非叶子节点，找到第一个大于该节点的节点替换
//
func (t *BinaryTree) Delete(val int) *TreeNode {
	if t.root == nil {
		return nil
	}

}

func (t *BinaryTree) MaxNode() *TreeNode {
	if t.root == nil {
		return nil
	}
	cur := t.root
	for {
		if cur.right == nil {
			return cur
		}
		cur = cur.right
	}
}

// MinNode 最小节点
func (t *BinaryTree) MinNode() *TreeNode {
	if t.root == nil {
		return nil
	}
	cur := t.root
	for {
		if cur.left == nil {
			return cur
		}
		cur = cur.left
	}
}
