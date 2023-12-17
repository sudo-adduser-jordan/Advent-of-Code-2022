package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(s string) (*bufio.Scanner, *os.File) {
	file, error := os.Open(s)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
}

func ReadInputFromFile() [][]string {
	return [][]string{}
}

func DoesExistFile() bool {
	return false
}

func DoesExistFolder() bool {
	return false
}
