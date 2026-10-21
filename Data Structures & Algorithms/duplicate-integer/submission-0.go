func hasDuplicate(nums []int) bool {
    mapping := map[int]int{}

    for _, val := range(nums){
        if _, found := mapping[val]; !found {
            mapping[val] = 1
        } else {
            return true
        }
    }
    return false
}
