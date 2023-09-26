package lds

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// sorting station algorithm
// https://ru.wikipedia.org/wiki/Алгоритм_сортировочной_станции

const (
	opPlus           = "+"
	opMinus          = "-"
	opMultiplication = "*"
	opDivision       = "/"
	opPower          = "^"
	//	opEquals         = "="
	funcSin      = "sin"
	funcCos      = "cos"
	openBracket  = "("
	closeBracket = ")"
)

func isOperator(token string) bool {
	switch token {
	case opPlus, opMinus, opMultiplication, opDivision, opPower:
		return true
	default:
		return false
	}
}

func isFunction(token string) bool {
	switch token {
	case funcSin, funcCos:
		return true
	default:
		return false
	}
}

func operatorPriority(operator string) int {
	switch operator {
	case opPower:
		return 0
	case opMultiplication, opDivision:
		return 1
	case opMinus, opPlus:
		return 2
	default:
		// whatever
		return 1337
	}
}

func isLeftAssociative(operator string) bool {
	switch operator {
	case opPlus, opMinus, opMultiplication, opDivision:
		return true
	case opPower:
		return false
	default:
		return false
	}
}

// everything is split with whitespace
// for ex: 2 + 2 * 2
// NOT 2+2*2
func SortingStation(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}
	stack := NewStack[string]()
	out := NewStack[string]()
	tokens := strings.Fields(string(data))
	// while there are still unprocessed tokens
	for _, token := range tokens {
		if isOperator(token) {
			// while there is operator op2 with higher priority or with same priority AND op1 is left-associative
			for !stack.Empty() && isOperator(stack.Peek()) &&
				((operatorPriority(stack.Peek()) < operatorPriority(token)) ||
					((operatorPriority(stack.Peek()) == operatorPriority(token)) && (isLeftAssociative(token)))) {
				operator := stack.Pop()
				out.Push(operator)
			}
			stack.Push(token)
		} else if isFunction(token) {
			stack.Push(token)
		} else if token == openBracket {
			stack.Push(token)
		} else if token == closeBracket {
			if stack.Empty() {
				return "", errors.New("expression is missing a closing bracket 1")
			}
			// until we don't see an open bracket at top of the stack
			for stack.Peek() != openBracket {
				operator := stack.Pop()
				out.Push(operator)
				if stack.Empty() {
					return "", errors.New("expression is missing a closing bracket 2")
				}
			}
			// get rid of the opening bracket
			stack.Pop()
			// if the function is present at the top of the stack - put it in the "out" stack
			if !stack.Empty() && isFunction(stack.Peek()) {
				function := stack.Pop()
				out.Push(function)
			}
		} else {
			// else token is a number - put it to the "out" stack
			out.Push(token)
		}
	}
	// while there are still tokens on the stack
	for !stack.Empty() {
		if stack.Peek() == openBracket {
			return "", errors.New("expression is missing a closing bracket 3")
		}
		operator := stack.Pop()
		out.Push(operator)
	}
	// now make an out stack into a string and write it to a file
	var builder strings.Builder
	for !out.Empty() {
		builder.WriteString(out.Pop())
		builder.WriteByte(' ')
	}
	return builder.String(), nil
}
