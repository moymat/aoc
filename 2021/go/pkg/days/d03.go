package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moymat/aoc2021/pkg/helpers"
)

type MostCommon int

const (
	Zero MostCommon = iota
	One
	Equal
)

func whichIsMoreCommon(inputs *[]string, idx int) MostCommon {
	ones := 0
	for _, input := range *inputs {
		if input[idx] == '1' {
			ones++
		}
	}
	zeros := len(*inputs) - ones
	if ones > zeros {
		return One
	}
	if ones < zeros {
		return Zero
	}
	return Equal
}

func updateLines(lines *[]string, compare rune, idx int) {
	var newLines []string
	for i, line := range *lines {
		if line[idx] != byte(compare) {
			newLines = append(newLines, (*lines)[i])
		}
	}
	*lines = newLines
}

func getRates(file string) (int64, int64) {
	lines := strings.Split(helpers.GetInput("d03", file), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	firstLine := strings.Split(lines[0], "")

	var gammaRates string
	for idx := range firstLine {
		if whichIsMoreCommon(&lines, idx) == One {
			gammaRates += "1"
		} else {
			gammaRates += "0"
		}
	}

	var epsilonRates string
	for _, num := range gammaRates {
		if num == '1' {
			epsilonRates += "0"
		} else {
			epsilonRates += "1"
		}
	}

	gammaNumber, err := strconv.ParseInt(gammaRates, 2, 64)
	helpers.CheckError(err)

	epsilonNumber, err := strconv.ParseInt(epsilonRates, 2, 64)
	helpers.CheckError(err)

	oxygenLines := append([]string{}, lines...)
	for idx := range firstLine {
		if len(oxygenLines) == 1 {
			break
		}

		if whichIsMoreCommon(&oxygenLines, idx) != Zero {
			updateLines(&oxygenLines, '0', idx)
		} else {
			updateLines(&oxygenLines, '1', idx)
		}

	}
	oxygenNumber, err := strconv.ParseInt(oxygenLines[0], 2, 64)
	helpers.CheckError(err)

	co2Lines := append([]string{}, lines...)
	for idx := range firstLine {
		if len(co2Lines) == 1 {
			break
		}

		if whichIsMoreCommon(&co2Lines, idx) == Zero {
			updateLines(&co2Lines, '0', idx)
		} else {
			updateLines(&co2Lines, '1', idx)
		}
	}
	co2Number, err := strconv.ParseInt(co2Lines[0], 2, 64)
	helpers.CheckError(err)

	fmt.Println(oxygenNumber, co2Number)

	return gammaNumber * epsilonNumber, oxygenNumber * co2Number
}

func RunD03(file string) {
	fmt.Println(getRates(file))
}
