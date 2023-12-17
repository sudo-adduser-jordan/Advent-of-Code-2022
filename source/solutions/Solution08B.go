package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
)

func PrintSolution08B(path string) {
	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	matrix := ParseGrid(scanner)
	score := utilities.ToString(FindPointsPartTwo(matrix))

	print("\t\t" + console.GREEN_TEXT)
	print(score)
	print(console.RESET_COLOR + "\n")
}

func FindPointsPartTwo(matrix [][]int) int {
	var score_slice []int
	var score int

	rows := len(matrix)
	for current_row := 1; current_row < rows-1; current_row++ {

		columns := len(matrix[current_row])
		for current_column := 1; current_column < columns-1; current_column++ {

			point := matrix[current_row][current_column]

			top_score := CheckTopPartTwo(matrix, current_row, current_column, point)
			bottom_score := CheckBottomPartTwo(matrix, current_row, current_column, point)
			left_score := CheckLeftPartTwo(matrix, current_row, current_column, point)
			right_score := CheckRightPartTwo(matrix, current_row, current_column, point)

			score = top_score * bottom_score * left_score * right_score
			score_slice = append(score_slice, score)
		}
	}
	return utilities.MaxIntSlice(score_slice)
}

func CheckTopPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for row := current_row - 1; row >= 0; row-- {
		count++
		if matrix[row][current_column] >= point {
			break
		}
	}
	return count
}

func CheckBottomPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for row := current_row + 1; row <= len(matrix)-1; row++ {
		count++
		if matrix[row][current_column] >= point {
			break
		}
	}
	return count
}

func CheckLeftPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for column := current_column - 1; column >= 0; column-- {
		count++
		if matrix[current_row][column] >= point {
			break
		}
	}
	return count
}

func CheckRightPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0

	for column := current_column + 1; column <= len(matrix)-1; column++ {
		count++
		if matrix[current_row][column] >= point {
			break
		}
	}
	return count
}
