func twoSum(nums []int, target int) []int {
	hm := map[int]int{}
	for i, v := range nums {
		need := target-v
		if _, ok := hm[need]; ok {
			return []int{hm[need], i}
		} else {
			hm[v] = i
		}
	}
	return []int{}
}
