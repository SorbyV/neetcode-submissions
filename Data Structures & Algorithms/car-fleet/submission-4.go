import (
	"slices"
)
func carFleet(target int, position []int, speed []int) int {
	hm := map[int]int{}
	for i := range(position) {
		hm[position[i]] = speed[i]
	}
	arrival := make([]float64, len(position))
	slices.Sort(position)

	for i := range(position) {
		pos := position[i]
		left := target-pos
		arrival[i] = float64(left)/float64(hm[pos])
	}

	fleets := len(arrival)
	lasttime := arrival[len(arrival)-1]
	for i := len(arrival)-2; i >= 0 ; i-- {
		currtime := arrival[i]
		if currtime <= lasttime {
			fleets -= 1
		} else {
			lasttime = currtime
		}
	}
	return fleets
}
