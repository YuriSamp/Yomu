package main

import (
	"os"
	"path/filepath"
	"tokei/output"
	"tokei/reader"
)

func main() {
	dir := "./test-folder"

	files, err := os.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	outputEntry := []reader.CodeInformation{}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		Output := reader.Read(filePath)
		outputEntry = append(outputEntry, Output )
	}

	output.OutputBuilder(outputEntry)
}
