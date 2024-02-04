/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/
package validParentheses

// runtime: 0ms, memory: 2.05MB
func isValid(s string) bool {
	if len(s)%2 == 1 || len(s) == 0 {
		return false
	}
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]rune, 0, len(s))

	for _, char := range s {
		if closing, ok := pairs[char]; ok {
			stack = append(stack, closing)
		} else if len(stack) == 0 || stack[len(stack)-1] != char {
			return false
		} else {
			// remove from stack
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
