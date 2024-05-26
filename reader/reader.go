package reader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CodeInformation struct {
	Extension string
	TotalLines int
	CommentedLines int
	CodeLines int
	BlankLines int
}

var extensionMap = map[string]string{
	"ts":   "Typescript",
	"js":   "Javascript",
	"json": "Json",
	"yaml": "Yaml",
	"go": "Go",
	"mod" : "Go",
	"md": "Markdown",
}

func Start(dir  string) {
	EmptyCodeInformationList := []CodeInformation{}
	markedPath := []string{}

	CodeInformationList, _ := readCurrDir(dir, EmptyCodeInformationList, markedPath)

	reduceInformation(CodeInformationList)
}

func readCurrDir(dir string, CodeInformationList []CodeInformation, markedPath []string) ([]CodeInformation,[]string ) {
		files, err := os.ReadDir(dir)

		if err != nil {
			panic(err)
		}

		for _ ,file := range files {

			if includes(markedPath, file.Name()) || strings.HasPrefix(file.Name(), ".") {
				continue
			}

			filePath := filepath.Join(dir, file.Name())

			if file.IsDir(){
				markedPath = append(markedPath, dir)
				CodeInformationList, markedPath = readCurrDir(filePath, CodeInformationList, markedPath)
			}

			if !file.IsDir() {
				extension := grabExtension(filePath)
				totalLines, commentedLines, codeLines, blankLines := readFileLines(filePath)

				
				info := CodeInformation{
            Extension:      extension,
            TotalLines:     totalLines,
            CommentedLines: commentedLines,
            CodeLines:      codeLines,
            BlankLines:     blankLines,
        }
        CodeInformationList = append(CodeInformationList, info)
			}
		}

		markedPath = append(markedPath, dir)
		return CodeInformationList, markedPath
}

func readFileLines(filePath string) (int, int, int, int)  {
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

	return len(lines), commentedLines, codeLines, blankLines
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

func includes(pathArray []string, word  string) bool {

	for _, path := range pathArray {
		if path == word {
			return true
		}
	}

	return false
}


func reduceInformation(informationList []CodeInformation) []CodeInformation {

	for _, info := range informationList {
		fmt.Println(info)
	}

	return informationList
}
