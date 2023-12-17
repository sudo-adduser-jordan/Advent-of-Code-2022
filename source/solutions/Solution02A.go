package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
)

func PrintSolution02A(path string) {

	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	var score int
	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "A X":
			score += 4
		case "A Y":
			score += 8
		case "A Z":
			score += 3

		case "B X":
			score += 1
		case "B Y":
			score += 5
		case "B Z":
			score += 9

		case "C X":
			score += 7
		case "C Y":
			score += 2
		case "C Z":
			score += 6
		default:
			println("error")
		}
	}

	print("\t\t" + console.GREEN_TEXT)
	print(score)
	print(console.RESET_COLOR + "\n")
}
