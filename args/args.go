package args

import (
	"fmt"
	"os"
	"Yomu/utils"
)

func HandleArgs() string {
	version := "0.0.1"

	args := os.Args[1:]

	if utils.Includes(args, "-v") {
		fmt.Printf("Yomu %s \n", version)
	}

	if len(args) == 0 {
		panic("Insert the directory you want to read")
	}

	dir := args[0]

	return dir
}
