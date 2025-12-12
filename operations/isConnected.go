package operations

func IsConnected(tetrominoes [][]string) bool {
	conect := 0
	result := true
	for _, tetetrominoe := range tetrominoes {

		for j, line := range tetetrominoe {
			for i, char := range line {
				if char == '#' {
					if i-1 >= 0 && line[i-1] == '#' {
						conect++
					}
					if i+1 < len(line) && line[i+1] == '#' {
						conect++
					}
					if j+1 < len(tetetrominoe) && tetetrominoe[j+1][i] == '#' {
						conect++
					}

					if j-1 >= 0 && tetetrominoe[j-1][i] == '#' {
						conect++
					}

				}
			}
		}
		if conect == 6 || conect == 8 {
			result = true
		}else{
			return false
		}
		conect = 0
	}

	return result
}
