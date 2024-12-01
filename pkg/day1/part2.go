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

	for _, list1Item := range list1 {
		for _, list2Item := range list2 {
			if list1Item == list2Item {
				frequencyMap[list1Item]++
			}
		}
	}

	sum := 0
	for k, v := range frequencyMap {
		sum += k * v
	}

	return sum, nil
}
