package solutions

import (
	"bufio"
	"bytes"
	"io"
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strings"
)

const ALPHA string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PrintSolution03A(path string) {
	priority := []int{}
	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		char := matchCharacterPartOne(line)
		x := numberValueOf(char)
		priority = append(priority, x)

	}
	result := sumArray(priority)

	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func numberValueOf(c rune) int {
	number := 0
	for index, value := range ALPHA {
		if c == value {
			number = index + 1
			break
		}
	}
	return number
}

func matchCharacterPartOne(line string) rune {
	var char rune
	midpoint := len(line) / 2

outerloop:
	for i := 0; i < midpoint; i++ {
		charOne := rune(line[i])

		for j := midpoint; j < len(line); j++ {
			charTwo := rune(line[j])

			if charOne == charTwo {
				char = charOne
				break outerloop
			}
		}
	}
	return char
}

func matchCharacterPartTwo(lineOne string, lineTwo string, lineThree string) rune {
	var c rune
	for i := 0; i < len(lineOne); i++ {
		c = rune(lineOne[i])
		if strings.ContainsRune(lineTwo, c) && strings.ContainsRune(lineThree, c) {
			return c
		}
	}
	return c
}

func LineCounter(r io.Reader) (int, error) {
	const lineBreak = '\n'
	count := 0
	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
