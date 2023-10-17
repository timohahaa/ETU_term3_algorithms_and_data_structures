package trees

import (
	"math/rand"
	"testing"
)

func TestBinaryTreeInsertSearchDelete(t *testing.T) {
	tree := NewBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	// search in an empty tree
	res := tree.Search(8)
	if res != nil {
		t.Errorf("the tree is empty, but search result is not nil: %+v", res)
	}

	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)

	// search again
	res = tree.Search(8)
	if res == nil {
		t.Errorf("the tree is not empty, but search result is nil: %+v", res)
	}
	if res != nil && res.val != 8 {
		t.Errorf("found wrong element: %+v, want: %d", res, 8)
	}

	// delete a random node
	val := 5 + rand.Intn(4)
	tree.Delete(val)
	// search again
	res = tree.Search(val)
	if res != nil {
		t.Errorf("found element, that was deleted: %+v", res)
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInorderTraversal(t *testing.T) {
	tree := NewBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	/*
	       5
	      / \
	     6   7
	    /
	   8
	*/
	want := []int{8, 6, 5, 7}
	res := tree.InorderTraversal()
	if !equal(want, res) {
		t.Errorf("got the wrong traversal: %v, want: %v", res, want)
	}
}

func TestPreorderTraversal(t *testing.T) {
	tree := NewBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	/*
	       5
	      / \
	     6   7
	    /
	   8
	*/
	want := []int{5, 6, 8, 7}
	res := tree.PreorderTraversal()
	if !equal(want, res) {
		t.Errorf("got the wrong traversal: %v, want: %v", res, want)
	}
}

func TestPostorderTraversal(t *testing.T) {
	tree := NewBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	/*
	       5
	      / \
	     6   7
	    /
	   8
	*/
	want := []int{8, 6, 7, 5}
	res := tree.PostorderTraversal()
	if !equal(want, res) {
		t.Errorf("got the wrong traversal: %v, want: %v", res, want)
	}
}
