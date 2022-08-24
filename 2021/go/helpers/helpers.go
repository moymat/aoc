package helpers

import (
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetInput(day string, file string) string {
	inputs, err := os.ReadFile("../inputs/" + day + "/" + file + ".txt")
	CheckError(err)
	return string(inputs)
}
