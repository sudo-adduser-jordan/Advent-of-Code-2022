package solutions

import (
	"bufio"
	"os"
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
	"strings"
	"unicode"
)

var Stacks []Stack

type Stack struct {
	crate []string
}

func (s *Stack) Peek() (element string) {
	return s.crate[len(s.crate)-1]
}

func (s *Stack) Push(element string) {
	s.crate = append(s.crate, element)
}

func (s *Stack) Pop() (element string) {
	element = s.crate[len(s.crate)-1]
	s.crate = s.crate[:len(s.crate)-1]
	return
}

func (s *Stack) PushToBottom(element string) {
	s.crate = append([]string{element}, s.crate...)
}

func (s *Stack) PushX(elements []string) {
	s.crate = append(s.crate, elements...)
}

func (s *Stack) PopX(n int) (elements []string) {
	elements = s.crate[len(s.crate)-n : len(s.crate)]
	s.crate = s.crate[:len(s.crate)-n]
	return
}

func PrintSolution05A(path string) {
	str := ""
	ParseInput(path)
	ParseInstructionsPartOne(path)
	for index := range Stacks {
		str += Stacks[index].Peek()
	}

	print("\t\t" + console.GREEN_TEXT)
	print(str)
	print(console.RESET_COLOR + "\n")
}

func ParseInput(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	str := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if ContainsOnlyNumbers(line) {
			break
		}
		str = line
	}
	str = strings.ReplaceAll(str, " ", "")
	slice := strings.Split(str, "")
	total_crates := len(slice)

	Stacks = make([]Stack, total_crates)

	file2, _ := os.Open(path)
	defer file.Close()
	scanner2 := bufio.NewScanner(file2)

	for scanner2.Scan() {
		line := scanner2.Text()
		if strings.Contains(line, "[") {
			for index, value := range line {
				if unicode.IsUpper(value) {
					Stacks[index/4].PushToBottom(string(value))
				}
			}
		} else {
			break
		}
	}
}

func ParseInstructionsPartOne(path string) {
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

			for i := 0; i < amount; i++ {
				element := Stacks[source-1].Pop()
				Stacks[destination-1].Push(element)
			}
		}
	}
}

func ContainsOnlyNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}
