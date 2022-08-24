package days

import (
	"fmt"
	"strconv"
	"strings"

	h "github.com/moymat/aoc2021/helpers"
)

func computeNbOfIncrease(file string) int {
	lines := strings.Split(h.GetInput("d01", file), "\n")
	total := 0

	for idx := range lines {
		if idx == 0 {
			continue
		}

		if idx == len(lines)-2 {
			break
		}

		previous, err := strconv.Atoi(strings.TrimSpace(lines[idx-1]))
		if err != nil {
			panic("not a number")
		}

		next, err := strconv.Atoi(strings.TrimSpace(lines[idx+2]))
		if err != nil {
			panic(("not a number"))
		}

		if next > previous {
			total += 1
		}
	}

	return total
}

func RunD01(file string) {
	fmt.Println(computeNbOfIncrease(file))
}
