package day2

import (
	"strconv"
	"strings"
)

func parse(in string) ([][]int, error) {
	lines := strings.Split(strings.TrimRight(in, "\n"), "\n")
	data := make([][]int, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, " ")
		level := make([]int, 0, len(parts))

		for _, part := range parts {
			v, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}

			level = append(level, v)
		}

		data = append(data, level)
	}

	return data, nil
}
