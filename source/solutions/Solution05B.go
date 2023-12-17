package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
	"strings"
)

func PrintSolution05B(path string) {
	str := ""
	ParseInput(path)
	ParseInstructionsPartTwo(path)
	for index := range Stacks {
		str += Stacks[index].Peek()
	}

	print("\t\t" + console.GREEN_TEXT)
	print(str)
	print(console.RESET_COLOR + "\n")
}

func ParseInstructionsPartTwo(path string) {
	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "m") {
			line := scanner.Text()
			array := strings.Fields(line)
			amount, _ := strconv.Atoi(array[1])
			source, _ := strconv.Atoi(array[3])
			destination, _ := strconv.Atoi(array[5])

			elements := Stacks[source-1].PopX(amount)
			Stacks[destination-1].PushX(elements)

		}
	}
}
