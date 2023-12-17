package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
)

// Sample
// R 5
// U 8
// L 8
// D 3
// R 17
// D 10
// L 25
// U 20

func PrintSolution09B(path string) {
	visited := map[Point]bool{}
	knots := [10]Point{}

	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		motions, _ := strconv.Atoi(line[2:])

		for motions > 0 {

			switch direction {
			case 'R':
				knots[0].x++
			case 'L':
				knots[0].x--
			case 'U':
				knots[0].y++
			case 'D':
				knots[0].y--
			}

			for i := 0; i < 9; i++ {
				knots[i+1] = MoveTailPart2(knots[i], knots[i+1])
			}

			visited[knots[9]] = true
			motions--
		}
	}

	result := utilities.ToString((len(visited)))
	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}

func MoveTailPart2(head Point, tail Point) Point {

	point := Point{head.x - tail.x, head.y - tail.y}

	switch point {
	case
		Point{-2, 1},
		Point{-1, 2},
		Point{0, 2},
		Point{1, 2},
		Point{2, 1},
		Point{2, 2},
		Point{-2, 2}:
		tail.y++
	}

	switch point {
	case
		Point{1, 2},
		Point{2, 1},
		Point{2, 0},
		Point{2, -1},
		Point{1, -2},
		Point{2, 2},
		Point{2, -2}:
		tail.x++
	}

	switch point {
	case
		Point{1, -2},
		Point{-1, -2},
		Point{0, -2},
		Point{2, -1},
		Point{2, -2},
		Point{-2, -2},
		Point{-2, -1}:
		tail.y--
	}

	switch point {
	case
		Point{-1, -2},
		Point{-1, 2},
		Point{-2, -2},
		Point{-2, -1},
		Point{-2, 0},
		Point{-2, 1},
		Point{-2, 2}:
		tail.x--
	}

	return tail
}
