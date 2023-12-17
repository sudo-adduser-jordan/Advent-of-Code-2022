package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
)

func PrintSolution01A(path string) {

	total := 0
	counter := 1
	calories := make(map[int]int)
	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			calories[counter] = total
			total = 0
			counter += 1
		} else {
			total += number
		}
	}

	result := 0
	for _, value := range calories {
		if value > result {
			result = value
		}
	}

	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}
