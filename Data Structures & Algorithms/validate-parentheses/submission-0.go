func isValid(s string) bool {
    stack := []byte{}
	for i := range len(s) {
		char := s[i]
		if char == '[' || char == '(' || char == '{' {
			stack = append(stack, char)
		fmt.Println(stack)
		} else {
			if char == ']' {
				if len(stack) == 0 {
					return false
				}
				if stack[len(stack)-1] != '[' {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			} else if char == ')' {
				if len(stack) == 0 {
					return false
				}
				if stack[len(stack)-1] != '(' {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			} else if char == '}' {
				if len(stack) == 0 {
					return false
				}
				if stack[len(stack)-1] != '{' {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	if len(stack) != 0 {

		return false
	}
	return true
}
