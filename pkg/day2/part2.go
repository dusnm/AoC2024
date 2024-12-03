package day2

import (
	"github.com/dusnm/AoC2024/pkg/input"
)

func Part2() (int, error) {
	reports, err := parse(input.Get("day2"))
	if err != nil {
		return 0, err
	}

	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		} else {
			for i := range report {
				newReport := removeLevel(report, i)
				if isReportSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports, nil
}
