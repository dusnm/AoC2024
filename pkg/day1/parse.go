package day1

import (
	"strconv"
	"strings"
)

type (
	Day1Data struct {
		List1 []int
		List2 []int
	}
)

func parse(in string) (Day1Data, error) {
	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)

	lines := strings.Split(in, "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		list1Part, err := strconv.Atoi(parts[0])
		if err != nil {
			return Day1Data{}, err
		}

		list2Part, err := strconv.Atoi(parts[1])
		if err != nil {
			return Day1Data{}, err
		}

		list1 = append(list1, list1Part)
		list2 = append(list2, list2Part)
	}

	return Day1Data{
		List1: list1,
		List2: list2,
	}, nil
}
