func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmap := map[byte]int{}
	for i := range(s){
		if _, found := hashmap[s[i]]; !found {
			hashmap[s[i]] = 1
		} else {
			hashmap[s[i]]++
		}
	}
	for j := range(t){
		if _, found := hashmap[t[j]]; !found {
			return false
		} else {
			hashmap[t[j]]--
		}
	}
	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}
		
		
	return true
}
