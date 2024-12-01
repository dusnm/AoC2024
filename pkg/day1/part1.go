package day1

import (
	"slices"

	"github.com/dusnm/AoC2024/pkg/input"
)

func Part1() (int, error) {
	in := input.Get("day1")
	data, err := parse(in)
	if err != nil {
		return 0, err
	}

	list1 := data.List1
	list2 := data.List2

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := range list1 {
		value1 := list1[i]
		value2 := list2[i]

		distance := value1 - value2
		if distance < 0 {
			distance = distance * -1
		}

		sum += distance
	}

	return sum, nil
}
