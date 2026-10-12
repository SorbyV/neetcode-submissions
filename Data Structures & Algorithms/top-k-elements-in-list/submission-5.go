func topKFrequent(nums []int, k int) []int {
	if len(nums)==k{
		return nums
	}
	hm := map[int]int{}
	for _, v := range nums {
		if _, ok := hm[v]; !ok {
			hm[v] = 1
		} else {
			hm[v] += 1
		}
	}
	ranks := make([][]int, len(nums)+1)
	for i := range ranks {
		ranks[i] = []int{}
	}
	ret := []int{}
	for k, v := range hm {
		ranks[v] = append(ranks[v], k)
	}
	i := len(ranks)-1
	for k > 0 {
		if len(ranks[i]) == 0 {
			i -= 1
			continue
		} else {
			for j := range ranks[i] {
				ret = append(ret, ranks[i][j])
			}
			k -= len(ranks[i])
			i -= 1
		}
	}
	return ret
}
