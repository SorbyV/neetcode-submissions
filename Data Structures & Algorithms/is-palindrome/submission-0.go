func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	fmt.Println(s)
	i := 0
	j := len(s)-1
	for i < j {
		if !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
			fmt.Println("i", string(s[i]))
			continue
		}
		if !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			fmt.Println("j", string(s[j]))
			j-=1
			continue
		}
		if s[i] != s[j] {
			return false
		} else {
			i+=1
			j-=1
			fmt.Println(string(s[i]), string(s[j]))
		}
	} 
	return true

}
