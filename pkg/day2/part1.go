package day2

import (
	"github.com/dusnm/AoC2024/pkg/input"
)

func Part1() (int, error) {
	in := input.Get("day2")
	reports, err := parse(in)
	if err != nil {
		return 0, err
	}

	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	return safeReports, nil
}
