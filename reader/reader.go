package reader

import (
	"os"
	"strings"
)

type CodeInformation struct {
	extension string
	lines int
	commentedLines int
	codeLines int
	blank int
}

var extensionMap = map[string]string{
	"ts":   "Typescript",
	"js":   "Javascript",
	"json": "Json",
	"yaml": "Yaml",
}

func Read(filePath string) CodeInformation {

	dat, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	extension := strings.Split(filePath, ".")
	extensionKey := extension[len(extension)-1]

	language, ok := extensionMap[extensionKey]

	if !ok {
		language = "Unknown"
	}

	blank, commentedLines, codeLines, lines := ReadFileLines(dat)

	return CodeInformation{blank: blank, codeLines: codeLines, commentedLines: commentedLines , extension: language, lines: lines}
}

func ReadFileLines(fileContentInBytes []byte) (int, int, int, int)  {
	fileContent := string(fileContentInBytes)

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

	return emptyLines, commentedLines, codeLines, len(lines)
}
