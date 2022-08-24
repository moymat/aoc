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
