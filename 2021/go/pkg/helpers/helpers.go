package helpers

import (
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetInput(day string, file string) string {
	dirname, _ := os.Getwd()
	path := strings.Replace(dirname, "2021/go", "2021/inputs/", -1)
	inputs, err := os.ReadFile(path + day + "/" + file + ".txt")
	CheckError(err)
	return string(inputs)
}

func IntSliceContains(slice *[]int, value int) bool {
	for _, val := range *slice {
		if val == value {
			return true
		}
	}
	return false
}

func MinMax(num1 int, num2 int) (int, int) {
	min := 0
	max := 0
	if num1 > num2 {
		min = num2
		max = num1
	} else {
		min = num1
		max = num2
	}
	return min, max
}
