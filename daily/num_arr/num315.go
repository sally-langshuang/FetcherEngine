package num_arr

func countSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	var tree, node *AVLTree
	for i := len(nums) - 1; i >= 0; i-- {
		tree, node = AddNode(tree, nums[i])
		ans[i] = CountLeft(tree, node)
	}
	return ans
}

type AVLTree struct {
	val int
	left, right *AVLTree
}
func (t * AVLTree) Constructor(val int) interface{}{
	return &AVLTree{val: val}
}
func (t * AVLTree) GetVal() int {
	return t.val
}
func (t * AVLTree) GetLeft()  interface{} {
	return t.left
}
func (t * AVLTree) SetLeft(left interface{}) {
	t.left = left.(*AVLTree)
}
func (t * AVLTree) GetRight() interface{} {
	return t.right
}
func (t * AVLTree) SetRight(right interface{}) {
	t.right = right.(*AVLTree)
}

func LLRole(root *AVLTree) *AVLTree {
	newRoot, rootLeft := root.left, root.left.right
	newRoot.right, root.left = root, rootLeft
	return newRoot
}

func RRRole(root *AVLTree) *AVLTree {
	newRoot, rootRight := root.right, root.right.left
	newRoot.left, root.right = root, rootRight
	return newRoot
}

func LRRole(root *AVLTree) *AVLTree  {
	newRoot, l, r := root.left.right, root.left.right.left, root.left.right.right
	newRoot.left, newRoot.right = root.left, root
	newRoot.left.right, newRoot.right.left = l, r
	return newRoot
}

func RLRole(root *AVLTree) *AVLTree  {
	newRoot, l, r := root.right.left, root.right.left.left, root.right.left.right
	newRoot.left, newRoot.right = root, root.right
	newRoot.left.right, newRoot.right.left = l, r
	return newRoot
}

func Balance(root *AVLTree)  *AVLTree{
	if root == nil || (root.left == nil && root.right == nil) {
		return root
	}
	newRoot := root
	root.left = Balance(root.left)
	root.right = Balance(root.right)
	if lDeep, rDeep := getDeepth(root.left), getDeepth(root.right); lDeep >= rDeep + 2 {
		if llDeep, lrDeep := getDeepth(root.left.left), getDeepth(root.left.right); llDeep >= lrDeep + 1 {
			newRoot = LLRole(root)
		} else {
			newRoot = LRRole(root)
		}
	} else if  rDeep >= lDeep + 2 {
		if rlDeep, rrDeep := getDeepth(root.right.left), getDeepth(root.right.right); rlDeep >= rrDeep + 1 {
			newRoot = RLRole(root)
		} else {
			newRoot = RRRole(root)
		}
	}
	return newRoot
}

func getDeepth(root *AVLTree) int {
	if root == nil {
		return 0
	}
	deepth := 1
	deepth +=  max(getDeepth(root.left), getDeepth(root.right))
	return deepth
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func CountLeft(root *AVLTree, node *AVLTree)  int {
	if root == node {
		return Count(root.left)
	}
	if root.val >= node.val {
		return CountLeft(root.left, node)
	} else {
		return Count(root.left) + 1 + CountLeft(root.right, node)
	}
}

func Count(root *AVLTree) int {
	if root == nil {
		return 0
	}
	return 1 + Count(root.left) + Count(root.right)
}

func AddNode(root *AVLTree, val int) (*AVLTree, *AVLTree) {
	node := &AVLTree{val: val}
	if root == nil {
		return node, node
	}
	for tmp := root; ;{
		var next *AVLTree
		if val > tmp.val {
			next = tmp.right
			if next == nil {
				tmp.right, node.right = node, next
				break
			}
		} else {
			next = tmp.left
			if next == nil {
				tmp.left, node.left = node, next
				break
			}
		}
		tmp = next
	}
	return Balance(root), node
}