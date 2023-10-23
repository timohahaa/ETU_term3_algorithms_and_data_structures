package trees

// the T type is litteraly anything
// provide a cmp(a, b T) int func when creating a tree, such that:
// if a < b, cmp(a,b) = -1 (or < 0)
// if a > b, cmp(a,b) = 1 (or > 0)
// if a = b, cmp(a,b) = 0

type BinaryTreeNode[T any] struct {
	Val   T
	right *BinaryTreeNode[T]
	left  *BinaryTreeNode[T]
}

func newBTNode[T any](val T) *BinaryTreeNode[T] {
	n := &BinaryTreeNode[T]{
		Val:   val,
		right: nil,
		left:  nil,
	}
	return n
}

type BinaryTree[T any] struct {
	root *BinaryTreeNode[T]
	cmp  func(a, b T) int
}

func NewBinaryTree[T any](cmp func(a, b T) int) *BinaryTree[T] {
	t := &BinaryTree[T]{
		root: nil,
		cmp:  cmp,
	}
	return t
}

func (t *BinaryTree[T]) SetRoot(root *BinaryTreeNode[T]) {
	t.root = root
}

func (t *BinaryTree[T]) inorder(arr *[]T, root *BinaryTreeNode[T]) {
	if root == nil {
		return
	}
	t.inorder(arr, root.left)
	*arr = append(*arr, root.Val)
	t.inorder(arr, root.right)
}

func (t *BinaryTree[T]) InorderTraversal() []T {
	traversalArr := make([]T, 0)
	t.inorder(&traversalArr, t.root)
	return traversalArr
}

func (t *BinaryTree[T]) preorder(arr *[]T, root *BinaryTreeNode[T]) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	t.preorder(arr, root.left)
	t.preorder(arr, root.right)
}

func (t *BinaryTree[T]) PreorderTraversal() []T {
	traversalArr := make([]T, 0)
	t.preorder(&traversalArr, t.root)
	return traversalArr
}

func (t *BinaryTree[T]) postorder(arr *[]T, root *BinaryTreeNode[T]) {
	if root == nil {
		return
	}
	t.postorder(arr, root.left)
	t.postorder(arr, root.right)
	*arr = append(*arr, root.Val)
}

func (t *BinaryTree[T]) PostorderTraversal() []T {
	traversalArr := []T{}
	t.postorder(&traversalArr, t.root)
	return traversalArr
}

func (t *BinaryTree[T]) LevelOrderTraversal() [][]T {
	if t.root == nil {
		return [][]T(nil)
	}
	ans := [][]T{}
	queue := []*BinaryTreeNode[T]{}
	queue = append(queue, t.root)
	for len(queue) != 0 {
		levelLen := len(queue)
		level := []T{}
		for i := 0; i < levelLen; i++ {
			node := queue[0]
			queue = queue[1:] // pop from queue
			level = append(level, node.Val)
			// add node's children to the queue
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		// finished current layer
		ans = append(ans, level)
	}
	return ans
}

func (t *BinaryTree[T]) height(root *BinaryTreeNode[T]) int {
	if root == nil {
		return 0
	}
	leftHeight := t.height(root.left)
	rightHeight := t.height(root.right)
	maxHeight := max(leftHeight, rightHeight)
	return maxHeight + 1
}

func (t *BinaryTree[T]) Height() int {
	return t.height(t.root)
}

func (t *BinaryTree[T]) search(root *BinaryTreeNode[T], elem T) *BinaryTreeNode[T] {
	if root == nil {
		return nil
	}

	if t.cmp(root.Val, elem) == 0 {
		return root
	}

	left := t.search(root.left, elem)
	right := t.search(root.right, elem)

	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func (t *BinaryTree[T]) Search(elem T) *BinaryTreeNode[T] {
	// returns nil if nothing found
	return t.search(t.root, elem)
}

// insert using bfs to the first nil-node, so the tree would be more balanced
func (t *BinaryTree[T]) Insert(elem T) {
	if t.root == nil {
		t.root = newBTNode[T](elem)
		return
	}
	// do the bfs-ing :)
	queue := []*BinaryTreeNode[T]{}
	queue = append(queue, t.root)
	for len(queue) != 0 {
		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			node := queue[0]
			queue = queue[1:] // pop from queue
			// add node's children to the queue
			if node.left != nil {
				queue = append(queue, node.left)
			} else {
				node.left = newBTNode[T](elem)
				return
			}
			if node.right != nil {
				queue = append(queue, node.right)
			} else {
				node.right = newBTNode[T](elem)
				return
			}
		}
	}
}

// replace with the deepest rightmost node then delete that node
// if elem does not exist, ignore
func (t *BinaryTree[T]) Delete(elem T) {
	elemPtr := t.Search(elem)
	if elemPtr == nil {
		return
	}
	// find the deepest rightmost node
	prev, cur := t.root, t.root
	for cur.right != nil {
		prev = cur
		cur = cur.right
	}
	// swap the data
	elemPtr.Val, cur.Val = cur.Val, elemPtr.Val
	// delete the node
	prev.right = nil
}
