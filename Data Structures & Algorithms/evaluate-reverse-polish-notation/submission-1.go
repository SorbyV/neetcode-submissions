func evalRPN(tokens []string) int {
	nums := []int{}

	for i := range len(tokens) {
		num, err := strconv.Atoi(tokens[i])
		if err == nil {
			nums = append(nums, num)
		} else if err != nil {
			num1 := nums[len(nums)-2]
			num2 := nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			if tokens[i] == "+" {
				nums = append(nums, num1+num2)
			} else if tokens[i] == "-" {
				nums = append(nums, num1-num2)
			} else if tokens[i] == "*" {
				nums = append(nums, num1*num2)
			} else if tokens[i] == "/" {
				nums = append(nums, num1/num2)
			}
		}
		
	}
	return nums[0]

}
