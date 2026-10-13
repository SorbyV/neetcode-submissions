func dailyTemperatures(temps []int) []int {
	res := make([]int, len(temps))
	stack := []int{0}
	// fmt.Println(len(temps))
	for i := 1; i < len(temps); i++ {
		curr := temps[i]
		// fmt.Println(res, stack, curr)
		for {
			if len(stack)==0 {
				// fmt.Println("stack emptied, adding: ", temps[i])
				stack = append(stack, i)
				break
			}
			top := temps[stack[len(stack)-1]]
			if top >= curr {
				// fmt.Println("Top of stack is greater than curr: ", top, curr)
				// increase score for top
				
				stack = append(stack, i)
				break
			}
			if top < curr {
				// fmt.Println("Top of stack is smaller/equal than curr: ", top, curr)

				res[stack[len(stack)-1]] += i-stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
		// fmt.Println("Now stack is: ", stack)
		
		if i == len(temps)-1 {
			for _, v := range stack {
			res[v] = 0
			}
		} 
	}

	
	return res
}