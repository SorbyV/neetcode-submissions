func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	i := 0
	j := len(s)-1
	for i < j {
		if !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			j-=1
			continue
		}
		if s[i] != s[j] {
			return false
		} else {
			i+=1
			j-=1
		}
	} 
	return true

}
