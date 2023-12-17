package solutions

import (
	"bufio"
	"fmt"
	"os"
	"source/cmd/packages/console"
)

func PrintSolution03B(path string) {
	priority := []int{}
	file, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	defer file.Close()

	len, error := LineCounter(file)
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	x := (len + 1) / 3
	for i := 0; i < x; i++ {

		scanner.Scan()
		lineOne := scanner.Text()
		scanner.Scan()
		lineTwo := scanner.Text()
		scanner.Scan()
		lineThree := scanner.Text()

		c := matchCharacterPartTwo(lineOne, lineTwo, lineThree)
		n := numberValueOf(c)
		priority = append(priority, n)
	}

	result := sumArray(priority)
	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}
