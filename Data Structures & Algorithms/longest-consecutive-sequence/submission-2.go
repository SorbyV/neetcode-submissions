func longestConsecutive(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := 1
	i := 0
	j := 1
	max := 1
	for j < len(nums) {

		if nums[i] == nums[j] {
			i = j
			j = i+1
		}else if nums[i]+1 == nums[j]{
			l += 1
			if l > max {
				max = l
			}
			i += 1
			j += 1
		} else {
			l = 1
			i = j
			j = i+1
		}
	}
	return max
}
