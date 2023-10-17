package trees

import (
	"errors"
	"strconv"
	"strings"
)

/*
example: (8 (9 (5)) (1))
gives:
           8
          / \
         9   1
        /
       5
*/
// no input validation, suppose that every input is valid
func parse(bracketString string) (*BinaryTreeNode[int], error) {
	// first and last symbols are always ( and )
	bracketString = bracketString[1 : len(bracketString)-1]

	firstOpenBracketIdx := strings.Index(bracketString, "(")
	// if no bracket is present, then there is only number left
	if firstOpenBracketIdx == -1 {
		rootVal, _ := strconv.Atoi(bracketString)
		return newBTNode[int](rootVal), nil
	}

	rootNumberStr := bracketString[:firstOpenBracketIdx]
	rootVal, _ := strconv.Atoi(rootNumberStr)
	root := newBTNode[int](rootVal)

	// cut the number off WITHOUT THE BRACKET
	bracketString = bracketString[firstOpenBracketIdx:]

	closingBraketIndex := findRespectiveClosingBracketIndex(bracketString)
	if closingBraketIndex == -1 {
		return nil, errors.New("incorrect input string: wrong brackets")
	}

	// cut subtrees
	leftSubtree, rightSubtree := bracketString[:closingBraketIndex+1], bracketString[closingBraketIndex+1:]
	// catch empty subtrees here, instead of making another recursion call
	var err error
	if len(leftSubtree) != 0 {
		root.left, err = parse(leftSubtree)
		if err != nil {
			return nil, err
		}
	}
	if len(rightSubtree) != 0 {
		root.right, err = parse(rightSubtree)
		if err != nil {
			return nil, err
		}
	}
	return root, nil
}

func findRespectiveClosingBracketIndex(bracketString string) int {
	// first symbol is the bracker - cut it off
	// then return the index + 1 to account for this cut
	bracketString = bracketString[1:]
	//	fmt.Println(bracketString)
	bracketPairCount := 0
	for idx := range bracketString {
		if bracketString[idx] == '(' {
			bracketPairCount++
		} else if bracketString[idx] == ')' {
			bracketPairCount--
		}
		if bracketPairCount == -1 {
			return idx + 1
		}
	}

	if bracketPairCount == -1 {
		return len(bracketString) // or -1 + 1
	}
	// should never return -1, or input is invalid
	return -1
}

func BinaryTreeFromBrackets(bracketString string) (*BinaryTree[int], error) {
	bracketString = strings.ReplaceAll(bracketString, " ", "")
	root, err := parse(bracketString)
	if err != nil {
		return nil, err
	}
	tree := NewBinaryTree[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	tree.SetRoot(root)
	return tree, nil
}
