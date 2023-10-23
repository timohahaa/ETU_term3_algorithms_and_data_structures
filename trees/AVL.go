package trees

import "fmt"

// the T type is litteraly anything
// provide a cmp(a, b T) int func when creating a tree, such that:
// if a < b, cmp(a,b) = -1 (or < 0)
// if a > b, cmp(a,b) = 1 (or > 0)
// if a = b, cmp(a,b) = 0

type AVLNode[T any] struct {
	Val    T
	left   *AVLNode[T]
	right  *AVLNode[T]
	parent *AVLNode[T]
	bf     int
}

func newAVLNode[T any](val T) *AVLNode[T] {
	n := &AVLNode[T]{
		Val:    val,
		left:   nil,
		right:  nil,
		parent: nil,
		bf:     0,
	}
	return n
}

type AVLTree[T any] struct {
	root *AVLNode[T]
	cmp  func(a, b T) int
}

func NewAVLTree[T any](cmp func(a, b T) int) *AVLTree[T] {
	t := &AVLTree[T]{
		root: nil,
		cmp:  cmp,
	}
	return t
}

func (t *AVLTree[T]) leftRotate(X *AVLNode[T]) {
	Y := X.right
	X.right = Y.left
	if Y.left != nil {
		Y.left.parent = X
	}
	Y.parent = X.parent
	if X.parent == nil {
		t.root = Y
	} else if X == X.parent.left {
		X.parent.left = Y
	} else {
		X.parent.right = Y
	}
	Y.left = X
	X.parent = Y

	X.bf = X.bf - 1 - max(0, Y.bf)
	Y.bf = Y.bf - 1 + min(0, X.bf)
}

func (t *AVLTree[T]) rightRotate(X *AVLNode[T]) {
	Y := X.left
	X.left = Y.right
	if Y.right != nil {
		Y.right.parent = X
	}
	Y.parent = X.parent
	if X.parent == nil {
		t.root = Y
	} else if X == X.parent.left {
		X.parent.left = Y
	} else {
		X.parent.right = Y
	}
	Y.right = X
	X.parent = Y

	X.bf = X.bf + 1 - min(0, Y.bf)
	Y.bf = Y.bf + 1 + max(0, X.bf)
}

func (t *AVLTree[T]) rebalance(node *AVLNode[T]) {
	if node.bf > 0 {
		if node.right.bf < 0 {
			t.rightRotate(node.right)
			t.leftRotate(node)
		} else {
			t.leftRotate(node)
		}
	} else if node.bf < 0 {
		if node.left.bf > 0 {
			t.leftRotate(node.left)
			t.rightRotate(node)
		} else {
			t.rightRotate(node)
		}
	}
}

func (t *AVLTree[T]) updateBalance(node *AVLNode[T]) {
	if node.bf < -1 || node.bf > 1 {
		t.rebalance(node)
		return
	}

	if node.parent != nil {
		if node == node.parent.left {
			node.parent.bf -= 1
		}
		if node == node.parent.right {
			node.parent.bf += 1
		}
		if node.parent.bf != 0 {
			t.updateBalance(node.parent)
		}
	}
}

func (t *AVLTree[T]) Insert(val T) {
	// first do ordinary BST insert
	node := newAVLNode[T](val)
	var y *AVLNode[T]
	x := t.root

	for x != nil {
		y = x
		if t.cmp(node.Val, x.Val) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	// y is parent of x
	node.parent = y
	if y == nil {
		t.root = node
	} else if t.cmp(node.Val, y.Val) < 0 {
		y.left = node
	} else {
		y.right = node
	}

	t.updateBalance(node)
}

func (t *AVLTree[T]) minimum(node *AVLNode[T]) *AVLNode[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *AVLTree[T]) delete(node *AVLNode[T], val T) *AVLNode[T] {
	// search for the key
	if node == nil {
		return node
	} else if t.cmp(val, node.Val) < 0 {
		node.left = t.delete(node.left, val)
	} else if t.cmp(val, node.Val) > 0 {
		node.right = t.delete(node.right, val)
	} else {
		var temp *AVLNode[T]
		if node.left == nil {
			temp = node.right
			node = nil
			return temp
		} else if node.right == nil {
			temp = node.left
			node = nil
			return temp
		}

		temp = t.minimum(node.right)
		node.Val = temp.Val
		node.right = t.delete(node.right, temp.Val)
	}

	if node == nil {
		return node
	}
	t.updateBalance(node)
	return node
}

func (t *AVLTree[T]) Delete(val T) {
	t.delete(t.root, val)
	return
}

func (t *AVLTree[T]) LevelOrderTraversal() [][]T {
	if t.root == nil {
		return [][]T(nil)
	}
	ans := [][]T{}
	queue := []*AVLNode[T]{}
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

func (t *AVLTree[T]) inorder(arr *[]T, root *AVLNode[T]) {
	if root == nil {
		return
	}
	t.inorder(arr, root.left)
	*arr = append(*arr, root.Val)
	t.inorder(arr, root.right)
}

func (t *AVLTree[T]) InorderTraversal() []T {
	traversalArr := make([]T, 0)
	t.inorder(&traversalArr, t.root)
	return traversalArr
}

func (t *AVLTree[T]) preorder(arr *[]T, root *AVLNode[T]) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	t.preorder(arr, root.left)
	t.preorder(arr, root.right)
}

func (t *AVLTree[T]) PreorderTraversal() []T {
	traversalArr := make([]T, 0)
	t.preorder(&traversalArr, t.root)
	return traversalArr
}

func (t *AVLTree[T]) postorder(arr *[]T, root *AVLNode[T]) {
	if root == nil {
		return
	}
	t.postorder(arr, root.left)
	t.postorder(arr, root.right)
	*arr = append(*arr, root.Val)
}

func (t *AVLTree[T]) PostorderTraversal() []T {
	traversalArr := []T{}
	t.postorder(&traversalArr, t.root)
	return traversalArr
}
func (t *AVLTree[T]) print(root *AVLNode[T], indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "   "
		} else {
			fmt.Print("L----")
			indent += "|  "
		}
		fmt.Printf("%v\n", root.Val)
		t.print(root.left, indent, false)
		t.print(root.right, indent, true)
	}
}

func (t *AVLTree[T]) PrintTree() {
	t.print(t.root, "", true)
}
