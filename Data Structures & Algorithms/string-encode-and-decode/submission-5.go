type Solution struct{}

func pad(l int) string {
	s := strconv.Itoa(l)
	if len(s) == 1 {
		s = "00" + s
	} else if len(s) == 2 {
		s = "0" + s
	}
	return s
}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "-1"
	}
	str := ""
	for i := range(strs) {
		str += pad(len(strs[i]))
		str += strs[i]
	}
	fmt.Println(str)
	return str
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "-1" {
		return []string{}
	}
	i := 0
	strs := []string{}
	for {
		l, _ := strconv.Atoi(string(encoded[i:i+3]))
		fmt.Println(l, encoded[i+3:i+l+3])
		strs = append(strs, encoded[i+3:i+l+3])
		i = i+l+3
		if i == len(encoded) {
			return strs
		}
	}
	return strs
}
