package main

import (
	data_structures "balanceBrackets/data-structures"
	"fmt"
	"strings"
)

/*
Challenge: The algorithm must balance the brackets of a string.
- If you have an empty pair of brackets, it should be removed.
- If you have extra closed brackets, it should be removed.

Example:
Input: “[a][b]][c]”
Output: ”[a][b][c]”
*/

func main() {
	const (
		inputSample  = "[a][b]][c]"
		expectOutput = "[a][b][c]"
	)

	countFormattedResult := BalanceUsingCountFormatted(inputSample)

	fmt.Println("Output using 'CountFormatted' was the same as expect:", strings.EqualFold(expectOutput, countFormattedResult))
	if !strings.EqualFold(expectOutput, countFormattedResult) {
		fmt.Println("CountFormatted:", countFormattedResult)
	}

}

// BalanceUsingStack receives an input string and retrieves a balanced string. Big O(n)
func BalanceUsingStack(input string) string {
	var (
		result  []string
		partial string
		stack   data_structures.Stack
	)

	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			stack.Push(string(input[i]))
		} else if input[i] == ']' {
			_, _ = stack.Pop()
			if len(partial) > 0 {
				result = append(result, fmt.Sprintf("%s]", partial))
				partial = ""
			}
		} else {
			partial += string(input[i])
		}
	}

	return fmt.Sprintf("[%s", strings.Join(result, "["))
}

// BalanceUsingCount receives an input string and retrieves a balanced string. Big O(n) *(in case of Go, it should be faster than use Stack struct built in this file)
func BalanceUsingCount(input string) string {
	var (
		result  []string
		partial string
		count   int
	)

	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			count++
		} else if input[i] == ']' {
			if count > 0 {
				count--
				if len(partial) > 0 {
					result = append(result, fmt.Sprintf("%s]", partial))
					partial = ""
				}
			}
		} else {
			partial += string(input[i])
		}
	}

	return fmt.Sprintf("[%s", strings.Join(result, "["))
}

func BalanceUsingCountFormatted(input string) (result string) {
	var (
		partial string
		count   int
	)

	for _, letter := range input {
		switch letter {
		case '[':
			count++
		case ']':
			if count > 0 {
				count--
				if len(partial) > 0 {
					result += fmt.Sprintf("[%s]", partial)
					partial = ""
				}
			}
		default:
			partial += string(letter)
		}
	}

	return
}

// BalanceUsingReplaceAll receives an input string and retrieves a balanced string. Big O(n^2)
func BalanceUsingReplaceAll(input string) string {
	var (
		result []string
		data   = strings.Split(input, "[")
	)

	for i := 0; i < len(data); i++ {
		data[i] = strings.ReplaceAll(data[i], "]", "")
		if len(data[i]) > 0 {
			result = append(result, fmt.Sprintf("%s]", data[i]))
		}
	}

	return fmt.Sprintf("[%s", strings.Join(result, "["))
}
