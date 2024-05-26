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

	for _, information := range entry {
		
		fmt.Println(information)
	}

	fmt.Println(divisor)

	fmt.Println("Total")

	fmt.Println(divisor)
}
