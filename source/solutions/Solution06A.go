package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strings"
)

func PrintSolution06A(path string) {
	signal := ""
	marker := ""
	counter := 5
	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		signal = scanner.Text()
	}

	marker = signal[0:4]
	signal = signal[4:]

	for _, value := range signal {
		marker = marker[1:] + string(value)
		if ValidateMarkerPartOne(marker) {
			break
		}
		counter++
	}

	print("\t\t" + console.GREEN_TEXT)
	print(counter)
	print(console.RESET_COLOR + "\n")
}

func ValidateMarkerPartOne(s string) bool {

	if strings.Contains(string(s[1:4]), string(s[0])) ||
		strings.Contains(string(s[0])+s[2:4], string(s[1])) ||
		strings.Contains(string(s[0:2])+s[3:4], string(s[2])) {
		return false
	}

	return true
}
