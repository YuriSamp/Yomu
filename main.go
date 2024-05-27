package main

import (
	"Yomu/args"
	"Yomu/output"
	"Yomu/reader"
)

func main() {
	dir, read := args.HandleArgs()

	if read {
		info := reader.Start(dir)
		
		output.OutputBuilder(info)
	}
}
