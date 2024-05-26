package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./teste-data.ts")

	if err != nil {
		panic(err)
	}

	fileContent := string(dat)

	lines := strings.Split(fileContent, "\n")

	emptyLines := 0
	commentedLines :=0
	codeLines := 0

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "//") {
			commentedLines += 1
			continue
		}

		if len(line) == 0{
			emptyLines += 1
			continue
		}

		codeLines += 1
	}

	fmt.Println("Empty Lines", emptyLines)
	fmt.Println("Commented Lines", commentedLines)
	fmt.Println("Code Lines", codeLines)
}
