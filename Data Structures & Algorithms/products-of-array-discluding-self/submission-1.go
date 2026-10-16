func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	i := 1
	j := len(nums)-2
	left[0] = 1
	right[len(nums)-1] = 1
	for {
		if i >= len(nums) {
			break
		}
		left[i] = left[i-1]*nums[i-1]
		right[j] = right[j+1]*nums[j+1]
		i++
		j--
	}
	ret := []int{}
	for i := range(nums) {
		ret = append(ret, left[i]*right[i])
	}
	return ret
}
