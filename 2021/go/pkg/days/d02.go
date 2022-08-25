package days

import (
	"strconv"
	"strings"

	"github.com/moymat/aoc2021/pkg/helpers"
)

type Position struct {
	x   int
	y   int
	aim int
}

func getMoveInput(input string, aim int) Position {
	parts := strings.Split(input, " ")

	amount, err := strconv.Atoi(parts[1])
	helpers.CheckError(err)

	if parts[0] == "forward" {
		return Position{x: amount, y: amount * aim, aim: aim}
	}
	if parts[0] == "down" {
		return Position{x: 0, y: 0, aim: aim + amount}
	}
	return Position{x: 0, y: 0, aim: aim - amount}
}

func getPosition(file string) int {
	finalPosition := Position{0, 0, 0}
	for _, line := range strings.Split(helpers.GetInput("d02", file), "\n") {
		position := getMoveInput(strings.TrimSpace(line), finalPosition.aim)
		finalPosition.x += position.x
		finalPosition.y += position.y
		finalPosition.aim = position.aim
	}
	return finalPosition.x * finalPosition.y
}

func RunD02(file string) {
	println(getPosition(file))
}
