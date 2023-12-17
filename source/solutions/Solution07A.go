package solutions

import (
	"bufio"
	"math"
	"os"
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strings"
)

type Directory struct {
	Name   string
	Parent *Directory
	Child  map[string]*Directory
	Files  map[string]int
	Size   int
}

func PrintSolution07A(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := ParseInput07(scanner)
	result := SumDirectories(root)

	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}

func ParseInput07(scanner *bufio.Scanner) *Directory {

	root := &Directory{
		Name:  "root",
		Child: map[string]*Directory{},
	}

	iterator := root

	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, " ")
		switch array[0] {
		// Execute Command
		case "$":
			switch array[1] {
			// Change Directory
			case "cd":
				switch array[2] {
				case "..":
					iterator = iterator.Parent
				default:
					if _, ok := iterator.Child[array[2]]; !ok {
						iterator.Child[array[2]] = &Directory{
							Name:   array[2],
							Parent: iterator,
							Child:  map[string]*Directory{},
							Files:  map[string]int{}}
					}
					iterator = iterator.Child[array[2]]
				}
			// Skip
			case "ls":
				continue
			}
		// Add Directory
		case "dir":
			if _, ok := iterator.Child[array[1]]; !ok {
				iterator.Child[array[1]] = &Directory{
					Name:   array[1],
					Parent: iterator,
					Child:  map[string]*Directory{},
					Files:  map[string]int{}}
			}
		// Add Files
		default:
			iterator.Files[array[1]] = utilities.ToInt(array[0])
		}
	}

	FillSize(root)
	return root

}

func FillSize(iterator *Directory) int {
	size := 0

	for _, v := range iterator.Child {
		size += FillSize(v)
	}
	for _, v := range iterator.Files {
		size += v
	}

	iterator.Size = size
	return size
}

func SumDirectories(iterator *Directory) int {
	limit := 100000
	sum := 0

	if iterator.Size <= limit {
		sum += iterator.Size
	}
	for _, v := range iterator.Child {
		sum += SumDirectories(v)
	}
	return sum
}

func FindDirectory(iterator *Directory, target int) int {
	min_size := math.MaxInt64
	if iterator.Size >= target {
		min_size = utilities.MinInt(min_size, iterator.Size)
	}

	for _, Child := range iterator.Child {
		min_size = utilities.MinInt(min_size, FindDirectory(Child, target))
	}

	return min_size
}
