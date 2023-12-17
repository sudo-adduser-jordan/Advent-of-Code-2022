package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(s string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
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
