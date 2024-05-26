package output

import (
	"fmt"
	"tokei/reader"
)

const (
	divisor = "==============================================================================="
)

func OutputBuilder(entry []reader.CodeInformation) {
	fmt.Println(divisor)

	fmt.Println("Language            Files        Lines         Code     Comments       Blanks")

	fmt.Println(divisor)

	totalLines := 0
	totalCodeLines := 0 
	totalCommentedLines := 0
	totalBlankLines := 0

	for _, information := range entry {
		fmt.Printf("%-12s %12d %12d %12d %12d %12d \n",
		information.Extension, 
		0,
		information.TotalLines, 
		information.CodeLines, 
		information.CommentedLines, 
		information.BlankLines)

		totalLines += information.TotalLines
		totalCommentedLines += information.CodeLines
		totalCodeLines += information.CommentedLines
		totalBlankLines += information.BlankLines
	}

	fmt.Println(divisor)

		fmt.Printf("%-12s %12d %12d %12d %12d %12d \n",
		"Total", 
		0,
		totalLines, 
		totalCodeLines, 
		totalCommentedLines, 
		totalBlankLines)

	fmt.Println(divisor)
}
