package solutions

import (
	"sort"
	"source/cmd/packages/console"
	"strings"
)

func PrintSolution11A(path string) {
	var monkeys []Monkey
	if strings.Contains(path, "Sample") {
		monkeys = MonkeyInputExample()
	} else {
		monkeys = MonkeyInput()
	}

	inspected := make([]int, len(monkeys))
	for round := 1; round <= 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspectedItem := monkey.operation(item) / 3
				recievingMonkey := monkey.test(inspectedItem)
				monkeys[recievingMonkey].items =
					append(monkeys[recievingMonkey].items, inspectedItem)
			}
			inspected[i] = inspected[i] + len(monkey.items)
			monkeys[i].items = []int{}
		}

	}

	sort.Ints(inspected)
	level := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	print("\t\t" + console.GREEN_TEXT)
	print(level)
	print(console.RESET_COLOR + "\n")
}
