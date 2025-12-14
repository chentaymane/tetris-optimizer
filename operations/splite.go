package operations

import (
	"strings"
)

func Split(input []byte) [][]string {
	temp := strings.ReplaceAll(string(input), "\r\n", "\n")

	splite := strings.Split(strings.TrimSpace(string(temp)), "\n")

	tetrominoes := [][]string{}
	tetrominoe := []string{}
	i := 0
	j := 0
	for i < len(splite) {

		if j < 4 {
			tetrominoe = append(tetrominoe, splite[i])
		} else {
			j = 0
			tetrominoes = append(tetrominoes, tetrominoe)
			tetrominoe = []string{}
			i++
			continue
		}
		j++
		i++
	}
	if tetrominoe != nil {
		tetrominoes = append(tetrominoes, tetrominoe)
	}
	return tetrominoes
}
