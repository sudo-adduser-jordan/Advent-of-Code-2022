package solutions

import (
	"sort"
	"source/cmd/packages/console"
	"strings"
)

func PrintSolution11B(path string) {
	var monkeys []Monkey
	if strings.Contains(path, "Sample") {
		monkeys = MonkeyInputExample()
	} else {
		monkeys = MonkeyInput()
	}

	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= m.divisor
	}

	inspected := make([]int, len(monkeys))
	for round := 1; round <= 10000; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspectedItem := monkey.operation(item)
				inspectedItem %= commonDivisor

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
