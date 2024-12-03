package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dusnm/AoC2024/pkg/day1"
	"github.com/dusnm/AoC2024/pkg/day2"
)

var (
	dayArg  uint
	partArg uint
)

func init() {
	flag.UintVar(
		&dayArg,
		"day",
		1,
		"select solutions for a given day: 1-25",
	)

	flag.UintVar(
		&partArg,
		"part",
		1,
		"select solution part: 1-2",
	)

	flag.Parse()
}

func main() {
	if dayArg < 1 || dayArg > 25 {
		fmt.Fprintf(os.Stderr, "day number out of range: %d\n", dayArg)
		os.Exit(-1)
	}

	if partArg != 1 && partArg != 2 {
		fmt.Fprintf(os.Stderr, "part number out of range: %d\n", partArg)
		os.Exit(-1)
	}

	days := [][]func() (int, error){
		{day1.Part1, day1.Part2},
		{day2.Part1, day2.Part2},
	}

	dayIndex := int(dayArg) - 1
	partIndex := int(partArg) - 1

	if len(days)-1 < dayIndex {
		fmt.Fprintf(os.Stderr, "no solution for day %d\n", dayArg)
		os.Exit(-1)
	}

	solutionFunc := days[dayIndex][partIndex]

	solution, err := solutionFunc()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}
