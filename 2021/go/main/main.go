package main

import (
	"fmt"
	"os"
	"strconv"

	d "github.com/moymat/aoc2021/days"
	h "github.com/moymat/aoc2021/helpers"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Need day number")
	}

	day, err := strconv.Atoi(os.Args[1])
	h.CheckError(err)

	file := "main"
	if len(os.Args) > 2 && os.Args[2] == "test" {
		file = "test"
	}

	fmt.Println("Running day " + os.Args[1] + " with " + file + " inputs")

	if day == 1 {
		d.RunD01(file)
	} else if day == 2 {
		d.RunD02(file)
	} else if day == 3 {
		d.RunD03(file)
	} else if day == 4 {
		d.RunD04(file)
	} else {
		panic("wrong day")
	}
}
