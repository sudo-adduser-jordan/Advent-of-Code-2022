package solutions

import (
	"source/cmd/packages/console"
	"source/cmd/packages/utilities"
	"strconv"
)

type Point struct {
	x int
	y int
}

func PrintSolution09A(path string) {

	head := Point{0, 0}
	tail := Point{0, 0}
	visited := map[Point]bool{}

	scanner, file := utilities.ScanFile(path)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		motions, _ := strconv.Atoi(line[2:])

		for motions > 0 {

			switch direction {
			case 'R':
				head.x++
			case 'L':
				head.x--
			case 'U':
				head.y++
			case 'D':
				head.y--
			}

			MoveTail(head, &tail)
			visited[tail] = true
			motions--
		}
	}

	result := utilities.ToString((len(visited)))
	print("\t\t" + console.GREEN_TEXT)
	print(result)
	print(console.RESET_COLOR + "\n")
}

func MoveTail(head Point, tail *Point) {

	if head.x-tail.x > 1 {
		tail.x++
		tail.y = head.y
	}

	if head.x-tail.x < -1 {
		tail.x--
		tail.y = head.y

	}

	if head.y-tail.y > 1 {
		tail.y++
		tail.x = head.x
	}

	if head.y-tail.y < -1 {
		tail.y--
		tail.x = head.x
	}

}
