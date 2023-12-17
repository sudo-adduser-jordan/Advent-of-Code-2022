package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
	"strings"
)

func PrintSolution04B(path string) {
	var total int
	scanner, file := utilities.ScanFile(path)
	defer file.Close()
	for scanner.Scan() {
		line := scanner.Text()

		array := strings.Split(line, ",")

		stringOne := array[0]
		stringTwo := array[1]

		arrayOne := strings.Split(stringOne, "-")
		arrayTwo := strings.Split(stringTwo, "-")

		rangeOneStart, _ := strconv.Atoi(arrayTwo[0])
		rangeOneEnd, _ := strconv.Atoi(arrayTwo[1])
		rangeTwoStart, _ := strconv.Atoi(arrayOne[0])
		rangeTwoEnd, _ := strconv.Atoi(arrayOne[1])

		if rangeOneStart <= rangeTwoEnd && rangeOneEnd >= rangeTwoStart ||
			rangeTwoStart <= rangeOneEnd && rangeTwoEnd >= rangeOneStart {
			total++
		}
	}

	print("\t\t" + console.GREEN_TEXT)
	print(total)
	print(console.RESET_COLOR + "\n")
}
