package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
)

func PrintSolution10A(path string) {
	cycles := 1
	registir := 1
	signal := map[int]int{}

	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		code := line[:4]

		switch code {
		case "noop":
			cycles++
			signal[cycles] = cycles * registir
		case "addx":
			value, _ := strconv.Atoi(line[5:])
			cycles++
			signal[cycles] = cycles * registir
			cycles++
			registir += value
			signal[cycles] = cycles * registir
		}
	}

	result :=
		signal[20] +
			signal[60] +
			signal[100] +
			signal[140] +
			signal[180] +
			signal[220]

	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}

func PrintCRT(spritePosition int, cycle int) {
	if cycle%40 == 0 {
		println()
	}

	if spritePosition-1 == cycle%40 ||
		spritePosition == cycle%40 ||
		spritePosition+1 == cycle%40 {
		print(console.RED_TEXT + "#" + console.RESET_COLOR)
	} else {
		print(".")
	}
}
