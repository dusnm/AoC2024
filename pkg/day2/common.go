package day2

import (
	"math"
)

func removeLevel(report []int, index int) []int {
	newReport := make([]int, 0, len(report)-1)
	for i, level := range report {
		if i != index {
			newReport = append(newReport, level)
		}
	}

	return newReport
}

func isReportSafe(report []int) bool {
	safe := true
	previousTrend := 0
	for i := 1; i < len(report); i++ {
		level := report[i]
		previousLevel := report[i-1]

		diff := level - previousLevel
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			safe = false
			break
		}

		trend := diff / absDiff
		if previousTrend != 0 && (previousTrend < 0) != (trend < 0) {
			safe = false
			break
		}

		previousTrend = trend
	}

	return safe
}
