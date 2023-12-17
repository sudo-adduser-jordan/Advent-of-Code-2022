package solutions

import (
	"source/cmd/packages/utilities"
	"strconv"
)

func PrintSolution10B(path string) {
	cycle := 0
	registir := 1

	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		code := line[:4]

		switch code {
		case "noop":
			PrintCRT(registir, cycle)
			cycle++
		case "addx":
			value, _ := strconv.Atoi(line[5:])
			PrintCRT(registir, cycle)
			cycle++
			PrintCRT(registir, cycle)
			cycle++
			registir += value
		}
	}
	println()
}
