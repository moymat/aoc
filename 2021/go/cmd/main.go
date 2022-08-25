package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/moymat/aoc2021/pkg/days"
	"github.com/moymat/aoc2021/pkg/helpers"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Need day number")
	}

	day, err := strconv.Atoi(os.Args[1])
	helpers.CheckError(err)

	file := "main"
	if len(os.Args) > 2 && os.Args[2] == "test" {
		file = "test"
	}

	fmt.Println("Running day " + os.Args[1] + " with " + file + " inputs")

	if day == 1 {
		days.RunD01(file)
	} else if day == 2 {
		days.RunD02(file)
	} else if day == 3 {
		days.RunD03(file)
	} else if day == 4 {
		days.RunD04(file)
	} else if day == 5 {
		days.RunD05(file)
	} else {
		panic("wrong day")
	}
}
