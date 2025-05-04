package src

func ValidParenthesis(s string) bool {
	ht := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ht['('] || char == ht['{'] || char == ht['['] {
			if stack[len(stack)-1] == ht[char] {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}

	return true
}
