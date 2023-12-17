package solutions

import (
	"bufio"
	"os"
	"source/cmd/packages/console"
)

func PrintSolution07B(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := ParseInput07(scanner)
	total_space := 70000000
	target := 30000000
	folder_to_delete := target - (total_space - root.Size)
	result := FindDirectory(root, folder_to_delete)

	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}
