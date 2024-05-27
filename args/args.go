package args

import (
	"Yomu/utils"
	"fmt"
	"os"
)

func HandleArgs() (string, bool) {
	version := "0.0.1"

	args := os.Args[1:]

	if utils.Includes(args, "-v") || utils.Includes(args, "--version") {
		fmt.Printf("Yomu %s \n", version)
		return "", false
	}

	if len(args) == 0 {
		panic("Insert at least one argument")
	}

	dir := args[0]

	return dir, true
}
