package operations

func Checker(tetrominoes [][]string) bool {

	for _, tetrominoe := range tetrominoes {
		if len(tetrominoe) != 4 {
			return false
		}
		count := 0
		for _, line := range tetrominoe {

			if len(line) != 4 {
				return false
			}
			for _, char := range line {
				if char != '.' && char != '#' {

					return false
				}
				if char == '#' {
					count++
				}
			}
		}
		if count != 4 {
			return false
		}
	}
	return true

}
