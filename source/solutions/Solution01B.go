package solutions

import (
	"sort"
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
)

func PrintSolution01B(path string) {

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

	keys := make([]int, 0, len(calories))
	for k := range calories {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return calories[keys[i]] < calories[keys[j]]
	})

	result := 0
	for i := len(keys) - 3; i < len(keys); i++ {
		result += calories[keys[i]]
	}

	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}
