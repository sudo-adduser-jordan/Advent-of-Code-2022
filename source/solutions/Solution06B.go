package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strings"
)

func PrintSolution06B(path string) {
	signal := ""
	marker := ""
	counter := 15
	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		signal = scanner.Text()
	}

	marker = signal[0:14]
	signal = signal[14:]

	for _, value := range signal {
		marker = marker[1:] + string(value)

		if ValidateMarkerPartTwo(marker) {
			break
		}

		counter++
	}
	print("\t\t" + console.GREEN_TEXT)
	print(counter)
	print(console.RESET_COLOR + "\n")
}

func ValidateMarkerPartTwo(s string) bool {

	for i := 0; i < len(s); i++ {
		if strings.Contains(
			strings.Replace(s, string(s[i]), "", 1),
			string(s[i])) {
			return false
		}
	}
	return true
}
