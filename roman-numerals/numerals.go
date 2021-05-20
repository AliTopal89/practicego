package numerals

import (
	"fmt"
	"strings"
)

func ConvertToRoman(num int) string {

	var result strings.Builder

	for i := num; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
		fmt.Println(result.String())
	}

	return result.String()
}
