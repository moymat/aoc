package days

import (
	"fmt"
	"strconv"
	"strings"

	h "github.com/moymat/aoc2021/helpers"
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

func updateLines(lines *[]string, linesCopy *[]string, compare rune, idx int) {
	for i, line := range *lines {
		if line[idx] == byte(compare) {
			if i < len(*linesCopy)-1 {
				*linesCopy = append((*linesCopy)[:i], (*linesCopy)[i+1:]...)
			} else {
				*linesCopy = (*linesCopy)[:i]
			}
		}
	}
}

func getRates(file string) (int64, int64) {
	lines := strings.Split(h.GetInput("d03", file), "\n")
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
	h.CheckError(err)

	epsilonNumber, err := strconv.ParseInt(epsilonRates, 2, 64)
	h.CheckError(err)

	oxygenLines := append([]string{}, lines...)
	for idx := range firstLine {
		if len(oxygenLines) == 1 {
			break
		}

		if whichIsMoreCommon(&lines, idx) != Zero {
			updateLines(&lines, &oxygenLines, '0', idx)
		} else {
			updateLines(&lines, &oxygenLines, '1', idx)
		}

	}
	oxygenNumber, err := strconv.ParseInt(oxygenLines[0], 2, 64)
	h.CheckError(err)

	co2Lines := append([]string{}, lines...)
	for idx := range firstLine {
		if len(co2Lines) == 1 {
			break
		}

		if whichIsMoreCommon(&lines, idx) == Zero {
			updateLines(&lines, &co2Lines, '0', idx)
		} else {
			updateLines(&lines, &co2Lines, '1', idx)
		}
	}
	co2Number, err := strconv.ParseInt(co2Lines[0], 2, 64)
	h.CheckError(err)

	fmt.Println(oxygenNumber, co2Number)

	return gammaNumber * epsilonNumber, oxygenNumber * co2Number
}

func RunD03(file string) {
	fmt.Println(getRates(file))
}
