func isValidSudoku(board [][]byte) bool {
	rows := make([]map[int]int, 9)
	cols := make([]map[int]int, 9)
	group := make([]map[int]int, 9)
	for i := range 9 {
    rows[i] = make(map[int]int)
	cols[i] = make(map[int]int)
	group[i] = make(map[int]int)
}
	for i := 0; i < 9; i++ {
		 //could be reduced to 3 instead of 9
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j])
			if _, ok := rows[i][num]; ok {
				return false
			}
			rows[i][num]++
			if _, ok := cols[j][num]; ok {
				return false
			}
			cols[j][num]++
			switch {
				case i>=0 && i < 3 && j >=0 && j < 3:
					if _, ok := group[0][num]; ok {
						return false
					}
					group[0][num]++
				case i>=0 && i < 3 && j >=3 && j < 6:
					if _, ok := group[1][num]; ok {
						return false
					}
					group[1][num]++
				case i>=0 && i < 3 && j >=6 && j < 9:
					if _, ok := group[2][num]; ok {
						return false
					}
					group[2][num]++
				case i>=3 && i < 6 && j >=0 && j < 3:
					if _, ok := group[3][num]; ok {
						return false
					}
					group[3][num]++
				case i>=3 && i < 6 && j >=3 && j < 6:
					if _, ok := group[4][num]; ok {
						return false
					}
					group[4][num]++
				case i>=3 && i < 6 && j >=6 && j < 9:
					if _, ok := group[5][num]; ok {
						return false
					}
					group[5][num]++
				case i>=6 && i < 9 && j >=0 && j < 3:
					if _, ok := group[6][num]; ok {
						return false
					}
					group[6][num]++
				case i>=6 && i < 9 && j >=3 && j < 6:
					if _, ok := group[7][num]; ok {
						return false
					}
					group[7][num]++
				case i>=6 && i < 9 && j >=6 && j < 9:
					if _, ok := group[8][num]; ok {
						return false
					}
					group[8][num]++			}
				
			
		}
	}
	return true
}
