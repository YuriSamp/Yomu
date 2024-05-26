package reader

import (
	"os"
	"strings"
)

func processFile(filePath string) CodeInformation  {
	extension := grabExtension(filePath)
	
	fileContentInBytes, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fileContent := string(fileContentInBytes)

	lines := strings.Split(fileContent, "\n")

	blankLines := 0
	commentedLines :=0
	codeLines := 0

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "//") {
			commentedLines += 1
			continue
		}

		if len(line) == 0 {
			blankLines += 1
			continue
		}

		codeLines += 1
	}

	return CodeInformation{TotalLines: len(lines), CommentedLines:  commentedLines,CodeLines:  codeLines, BlankLines:  blankLines, Extension: extension, Files: 1 }
}

func grabExtension(fileName string ) string {
	extension := strings.Split(fileName, ".")
	extensionKey := extension[len(extension)-1]

	language, ok := extensionMap[extensionKey]

	if !ok {
		language = "Unknown"
	}

	return language
}
