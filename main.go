package main

import (
	"fmt"
	"os"
	"tokei/reader"
)

func main() {
	dir := handleArgs()

	reader.Start(dir)
	
	// outputEntry := []reader.CodeInformation{}
	// output.OutputBuilder(outputEntry)
}

func handleArgs() string {
	args := os.Args[1:]

	fmt.Println(args)

	dir := args[0]

	return dir
}
