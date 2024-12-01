package day1

import "github.com/dusnm/AoC2024/pkg/input"

func Part2() (int, error) {
	in := input.Get("day1")
	data, err := parse(in)
	if err != nil {
		return 0, err
	}

	list1 := data.List1
	list2 := data.List2

	frequencyMap := make(map[int]int, len(list1))
	for _, item := range list2 {
		frequencyMap[item]++
	}

	sum := 0
	for _, item := range list1 {
		v, ok := frequencyMap[item]
		if ok {
			sum += item * v
		}
	}

	return sum, nil
}
